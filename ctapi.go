package main

import (
	"context"
	"fmt"
	"net/url"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/cloudtruth/cloudtruth-go-client"
)

type CTApi struct {
	configuration         *cloudtruth.Configuration
	client                *cloudtruth.APIClient
	environment_ids_cache map[string]string
	project_ids_cache     map[string]string
	template_ids_cache    map[string]map[string]string
}

type Parameter struct {
	key    string
	value  string
	secret bool
}

type Template struct {
	key   string
	value func() string
}

func NewCTApi(api_key string, api_url string, user_agent string, debug bool) *CTApi {
	log.Debug("Creating new CTApi")

	ctapi := new(CTApi)
	ctapi.configuration = cloudtruth.NewConfiguration()
	ctapi.configuration.UserAgent = user_agent
	ctapi.template_ids_cache = make(map[string]map[string]string)
	u, err := url.Parse(api_url)
	if err != nil {
		log.Fatalf("Invalid url: %s, err: %v", api_url, err)
	}
	ctapi.configuration.Host = u.Host
	ctapi.configuration.Scheme = u.Scheme
	ctapi.configuration.Debug = debug

	ctapi.configuration.AddDefaultHeader("Authorization", fmt.Sprintf("Api-Key %s", api_key))

	ctapi.client = cloudtruth.NewAPIClient(ctapi.configuration)

	return ctapi
}

func (ctapi *CTApi) environment_ids() map[string]string {
	if ctapi.environment_ids_cache == nil {
		log.Debug("Fetching environments")

		resp, r, err := ctapi.client.EnvironmentsApi.EnvironmentsList(context.Background()).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.EnvironmentsList``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
			os.Exit(2)
		}
		// log.Printf("Response from `EnvironmentsApi.EnvironmentsList`: %v\n", resp)
		ctapi.environment_ids_cache = make(map[string]string)
		for _, p := range resp.Results {
			ctapi.environment_ids_cache[p.Name] = p.Id
		}
	}

	return ctapi.environment_ids_cache
}

func (ctapi *CTApi) environments() []string {
	id_map := ctapi.environment_ids()

	keys := make([]string, 0, len(id_map))
	for k := range id_map {
		keys = append(keys, k)
	}
	return keys
}

func (ctapi *CTApi) project_ids() map[string]string {
	if ctapi.project_ids_cache == nil {
		log.Debug("Fetching projects")

		resp, r, err := ctapi.client.ProjectsApi.ProjectsList(context.Background()).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsList``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
			os.Exit(2)
		}
		// log.Printf("Response from `ProjectsApi.ProjectsList`: %v\n", resp)
		ctapi.project_ids_cache = make(map[string]string)
		for _, p := range resp.Results {
			ctapi.project_ids_cache[p.Name] = p.Id
		}
	}

	return ctapi.project_ids_cache
}

func (ctapi *CTApi) projects() []string {
	id_map := ctapi.project_ids()

	keys := make([]string, 0, len(id_map))
	for k := range id_map {
		keys = append(keys, k)
	}
	return keys
}

func (ctapi *CTApi) parameters(project string, environment string, tag string) map[string]Parameter {
	log.Debug("Fetching parameters")

	env_id := ctapi.environment_ids()[environment]
	proj_id := ctapi.project_ids()[project]

	request := ctapi.client.ProjectsApi.ProjectsParametersList(context.Background(), proj_id)
	request = request.Environment(env_id)
	request = request.Evaluate(true)
	if tag != "" {
		request = request.Tag(tag)
	}

	resp, r, err := request.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsParametersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		os.Exit(2)
	}
	// log.Printf("Response from `ProjectsApi.ProjectsParametersList`: %v\n", resp)
	result := make(map[string]Parameter)
	for _, p := range resp.Results {
		var v cloudtruth.ParameterValuesValue
		for _, v = range p.Values {
			break
		}
		result[p.Name] = Parameter{p.Name, *v.Value.Get(), *p.Secret}
	}

	return result
}

func (ctapi *CTApi) template_ids(project string) map[string]string {
	return ctapi.template_ids_cache[project]
}

func (ctapi *CTApi) template(project string, environment string, tag string, template string) string {

	env_id := ctapi.environment_ids()[environment]
	proj_id := ctapi.project_ids()[project]
	tmpl_id := ctapi.template_ids(project)[template]
	log.Debugf("Fetching template %s/%s/%s", project, template, tmpl_id)

	request := ctapi.client.ProjectsApi.ProjectsTemplatesRetrieve(context.Background(), tmpl_id, proj_id)
	request = request.Environment(env_id)
	if tag != "" {
		request = request.Tag(tag)
	}

	resp, r, err := request.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsTemplatesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		os.Exit(2)
	}
	// log.Printf("Response from `ProjectsApi.ProjectsTemplatesRetrieve`: %v\n", resp)

	return *resp.Body
}

func (ctapi *CTApi) templates(project string, environment string, tag string) map[string]Template {
	log.Debug("Fetching templates")

	template_ids_cache := make(map[string]string)

	env_id := ctapi.environment_ids()[environment]
	proj_id := ctapi.project_ids()[project]

	request := ctapi.client.ProjectsApi.ProjectsTemplatesList(context.Background(), proj_id)
	request = request.Environment(env_id)
	request = request.Evaluate(false)
	if tag != "" {
		request = request.Tag(tag)
	}

	resp, r, err := request.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsTemplatesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		os.Exit(2)
	}
	// log.Printf("Response from `ProjectsApi.ProjectsTemplatesList`: %v\n", resp)
	result := make(map[string]Template)
	for _, t := range resp.Results {
		name := t.Name
		log.Debugf("Adding template reference '%s'", name)
		anon := func() string { return ctapi.template(project, environment, tag, name) }
		result[t.Name] = Template{name, anon}
		template_ids_cache[t.Name] = t.Id
	}
	ctapi.template_ids_cache[project] = template_ids_cache

	return result
}
