package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
	"github.com/yargevad/filepathx"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

type Options struct {
	Verbose          []bool   `short:"v" long:"verbose" description:"Show verbose debug information"`
	Version          bool     `long:"version" description:"Show version information"`
	Environment      string   `short:"e" long:"environment" description:"The cloudtruth environment" default:"default" env:"CLOUDTRUTH_ENVIRONMENT"`
	Project          string   `short:"p" long:"project" description:"The cloudtruth project" env:"CLOUDTRUTH_PROJECT"`
	Tag              string   `short:"t" long:"tag" description:"The environment tag to restrict values to" env:"CLOUDTRUTH_TAG"`
	ReferencePattern string   `short:"r" long:"reference-pattern" description:"The reference pattern (go fmt) to substitute with parameters" default:"<%s>" env:"CLOUDTRUTH_REFERENCE_PATTERN"`
	FilePattern      []string `short:"f" long:"file-pattern" description:"The file pattern (glob) to perform substitutions on" default:"*.y*ml" env:"CLOUDTRUTH_FILE_PATTERN" env-delim:","`
	ApiKey           string   `short:"a" long:"api-key" description:"The cloudtruth api key" env:"CLOUDTRUTH_API_KEY"`
	ApiUrl           string   `short:"u" long:"api-url" description:"The cloudtruth api url" env:"CLOUDTRUTH_API_URL" hidden:"true" default:"https://api.cloudtruth.io"`
}

func mergeArgoEnv(name string, defaultValue string, pluginConfig map[string]string) string {
	value := defaultValue

	// Sytem env - we get defaults from here, they are set via in an Secret mounted with envFrom in the patch applied to
	// the argocd-repo-server in our setup.sh
	v := os.Getenv(name)
	if v != "" {
		value = v
	}

	// Params set within the yaml file used to trigger auto discovery of the plugin
	v = pluginConfig[strings.Replace(strings.ToLower(name), "_", "-", -1)]
	if v != "" {
		value = v
	}

	// Env set for the plugin as part of its attachment to the argocd Application
	v = os.Getenv("ARGOCD_ENV_" + name)
	if v != "" {
		value = v
	}

	// Params set for the plugin as part of its attachment to the argocd Application
	v = os.Getenv("PARAM_" + name)
	if v != "" {
		value = v
	}

	return value
}

func readPluginConfig(filename string) map[string]string {
	log.Debugf("Parsing plugin config file '%s'", filename)
	data := make(map[string]string)
	yfile, err := os.ReadFile(filename)

	if err != nil {
		log.Warnf("Unable to read plugin config from file '%s': %s", filename, err)
		return data
	}

	err2 := yaml.Unmarshal(yfile, &data)

	if err2 != nil {
		log.Warnf("Could not parse yaml from plugin config file '%s': %s", filename, err2)
		return data
	}

	return data
}

func setLogLevel(level int) {
	switch level {
	case 2:
		log.SetLevel(log.DebugLevel)
	case 1:
		log.SetLevel(log.InfoLevel)
	case 0:
		log.SetLevel(log.WarnLevel)
	default:
		log.SetLevel(log.TraceLevel)
	}
}

// Processes given files to replace paramater references with values from cloudtruth
func main() {
	log.SetOutput(os.Stderr)

	var config Options
	p := flags.NewParser(&config, flags.Default)
	p.LongDescription = "Processes given files to replace paramater references with values from cloudtruth."
	_, err := p.Parse()

	if config.Version {
		fmt.Printf("argocd-cloudtruth-plugin %s, commit %s, built at %s by %s\n", version, commit, date, builtBy)
		os.Exit(1)
	}

	if err != nil {
		os.Exit(1)
	}

	pluginConfig := readPluginConfig(".argocd-cloudtruth-plugin")

	log_level := mergeArgoEnv("CLOUDTRUTH_LOG_LEVEL", strconv.Itoa(len(config.Verbose)), pluginConfig)
	log_level_int, err := strconv.Atoi(log_level)
	if err != nil {
		fmt.Printf("Invalid log level '%s'\n", log_level)
		os.Exit(1)
	}
	setLogLevel(log_level_int)

	config.ApiUrl = mergeArgoEnv("CLOUDTRUTH_API_URL", config.ApiUrl, pluginConfig)
	config.ApiKey = mergeArgoEnv("CLOUDTRUTH_API_KEY", config.ApiKey, pluginConfig)
	config.Environment = mergeArgoEnv("CLOUDTRUTH_ENVIRONMENT", config.Environment, pluginConfig)
	config.Project = mergeArgoEnv("CLOUDTRUTH_PROJECT", config.Project, pluginConfig)
	config.Tag = mergeArgoEnv("CLOUDTRUTH_TAG", config.Tag, pluginConfig)
	config.ReferencePattern = mergeArgoEnv("CLOUDTRUTH_REFERENCE_PATTERN", config.ReferencePattern, pluginConfig)
	config.FilePattern = strings.Split(mergeArgoEnv("CLOUDTRUTH_FILE_PATTERN", strings.Join(config.FilePattern, ","), pluginConfig), ",")

	if config.ApiKey == "" {
		fmt.Printf("An api key is required\n")
		os.Exit(1)
	}

	if config.Project == "" {
		fmt.Printf("A project is required\n")
		os.Exit(1)
	}

	log.Debug("ApiUrl: ", config.ApiUrl)
	log.Debug("Environment: ", config.Environment)
	log.Debug("Project: ", config.Project)
	log.Trace("ALL Config: ", config)

	// TODO: allow user to specify project and/or environment in the replacement pattern, e.g. <ENV:PROJ:PARAM>
	// TODO: scan files to figure out which ones have a pattern to be replaced rather than replacing against all files
	// TODO: support templates - how do we map a template to a file, or just have a file that only contains the template?
	ctapi := NewCTApi(config.ApiKey, config.ApiUrl, fmt.Sprintf("argocd-cloudtruth-plugin/%s/%s/go", version, commit), log_level_int >= 4)
	ctproject := NewCTProject(ctapi, config.Project, config.Environment, config.Tag)

	err = applyTransformations(config.FilePattern, config.ReferencePattern, ctproject)
	if err != nil {
		log.Fatalf("Failed to apply transformations: %v", err)
	}
}

func applyTransformations(filePattern []string, referencePattern string, ctproject *CTProject) error {
	first := true
	for _, pattern := range filePattern {
		log.Info("Processing pattern: ", pattern)

		matches, err := filepathx.Glob(pattern)
		if err != nil {
			log.Errorf("Failed to match pattern: %v", err)
			return err
		}

		for _, path := range matches {
			log.Info("Processing file pattern match: ", path)

			if !first {
				fmt.Print("\n\n---\n\n")
			}
			result, e := fileReplace(path, referencePattern, ctproject)
			if e != nil {
				return e
			}
			fmt.Print(result)
			first = false
		}
	}
	return nil
}

func fileReplace(path string, pattern string, ctproject *CTProject) (string, error) {
	originalContents, err := os.ReadFile(path)
	if err != nil {
		log.Errorf("Unable to read file: %v", err)
		return "", err
	}

	replacedContents := string(originalContents)

	// Kinda hacky, but no option due to use of golang tring format matches for replacement (<%s>)
	// Should probably change the interface on the cli to use regexes (?), or maybe even use full liquid templates for plugin
	templatesPresent := false
	paramsPresent := false
	re := regexp.MustCompile(strings.Replace(pattern, "%s", "(.+?)", -1))
	refs := re.FindAllStringSubmatch(replacedContents, -1)

	for _, ref := range refs {
		if len(ref) > 1 && strings.HasPrefix(ref[1], "templates.") {
			templatesPresent = true
		} else {
			paramsPresent = true
		}
	}

	if paramsPresent {
		for k, v := range ctproject.Parameters() {
			pattern := fmt.Sprintf(pattern, k)
			replacedContents = strings.Replace(replacedContents, pattern, v.value, -1)
		}
	}

	if templatesPresent {
		for k, v := range ctproject.Templates() {
			pattern := fmt.Sprintf(pattern, "templates."+k)
			if strings.Contains(replacedContents, pattern) {
				log.Debugf("Replacing template '%s'", pattern)
				template_value := v.Get()
				replacedContents = strings.Replace(replacedContents, pattern, template_value, -1)
			}
		}
	}

	log.Tracef("**** Start Transformation result of '%s':\n%s\n**** End Transformation result of '%s'\n", path, replacedContents, path)
	return replacedContents, nil
}
