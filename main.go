package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
	"github.com/yargevad/filepathx"
)

var config struct {
	Verbose          []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
	Environment      string `short:"e" long:"environment" description:"The cloudtruth environment" env:"CLOUDTRUTH_ENVIRONMENT"`
	Project          string `short:"p" long:"project" description:"The cloudtruth project" env:"CLOUDTRUTH_PROJECT" required:"true"`
	Tag              string `short:"t" long:"tag" description:"The environment tag to restrict values to" env:"CLOUDTRUTH_TAG"`
	ReferencePattern string `short:"r" long:"reference-pattern" description:"The reference pattern (go fmt) to substitute with parameters" default:"<%s>" env:"CLOUDTRUTH_REFERENCE_PATTERN"`
	ApiKey           string `short:"a" long:"api-key" description:"The cloudtruth api key" env:"CLOUDTRUTH_API_KEY" required:"true"`
	ApiUrl           string `short:"u" long:"api-url" description:"The cloudtruth api url" env:"CLOUDTRUTH_API_URL" hidden:"true" default:"https://api.cloudtruth.io"`
	Positional       struct {
		FilePatterns []string `positional-arg-name:"FILEGLOB" required:"true"`
	} `positional-args:"true"`
}

//Processes given files to replace paramater references with values from cloudtruth

func main() {

	p := flags.NewParser(&config, flags.Default)
	p.LongDescription = "Processes given files to replace paramater references with values from cloudtruth."
	_, err := p.Parse()
	if err != nil {
		os.Exit(1)
	}

	log.SetOutput(os.Stderr)
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

	log.Debug("ApiUrl: ", config.ApiUrl)
	log.Debug("Environment: ", config.Environment)
	log.Debug("Project: ", config.Project)

	ctapi := NewCTApi(config.ApiKey, config.ApiUrl)

	// TODO: allow user to specify project and/or environment in the replacement pattern, e.g. <ENV:PROJ:PARAM>
	params := ctapi.parameters(config.Project, config.Environment, config.Tag)

	// TODO: scan files to figure out which ones have a pattern to be replaced rather than replacing against all files
	first := true
	for _, pattern := range config.Positional.FilePatterns {
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
