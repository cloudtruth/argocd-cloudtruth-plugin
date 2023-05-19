package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

var config struct {
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
	yfile, err := ioutil.ReadFile(filename)

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

// Processes given files to replace paramater references with values from cloudtruth
func main() {
	log.SetOutput(os.Stderr)

	p := flags.NewParser(&config, flags.Default)
	p.LongDescription = "Processes given files to replace paramater references with values from cloudtruth."
	_, err := p.Parse()

	switch len(config.Verbose) {
	case 2:
		log.SetLevel(log.DebugLevel)
	case 1:
		log.SetLevel(log.InfoLevel)
	case 0:
		log.SetLevel(log.WarnLevel)
	default:
		log.SetLevel(log.TraceLevel)
	}

	if config.Version {
		fmt.Printf("argocd-cloudtruth-plugin %s, commit %s, built at %s by %s\n", version, commit, date, builtBy)
		os.Exit(1)
	}

	if err != nil {
		os.Exit(1)
	}

	pluginConfig := readPluginConfig(".argocd-cloudtruth-plugin")
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

	ctapi := NewCTApi(config.ApiKey, config.ApiUrl, fmt.Sprintf("argocd-cloudtruth-plugin/%s/%s/go", version, commit))

	// TODO: allow user to specify project and/or environment in the replacement pattern, e.g. <ENV:PROJ:PARAM>
	params := ctapi.parameters(config.Project, config.Environment, config.Tag)

	// TODO: scan files to figure out which ones have a pattern to be replaced rather than replacing against all files
	first := true
	for _, pattern := range config.FilePattern {
		log.Info("Processing pattern: ", pattern)

		matches, err := filepathx.Glob(pattern)
		if err != nil {
			log.Fatal(err)
		}

		for _, path := range matches {
			log.Info("Processing file pattern match: ", path)

			if !first {
				fmt.Print("\n\n---\n\n")
			}
			fmt.Print(fileReplace(path, config.ReferencePattern, params))
			first = false
		}
	}
}

func fileReplace(path string, pattern string, parameters map[string]Parameter) string {
	originalContents, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}

	replacedContents := string(originalContents)
	for k, v := range parameters {
		pattern := fmt.Sprintf(pattern, k)
		replacedContents = strings.Replace(replacedContents, pattern, v.value, -1)
	}

	return replacedContents
}
