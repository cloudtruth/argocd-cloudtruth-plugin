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
	configuration *cloudtruth.Configuration
	client        *cloudtruth.APIClient
	environments  map[string]cloudtruth.Environment
	projects      map[string]cloudtruth.Project
}

type CTProject struct {
	ctapi            *CTApi
	environment_name string
	environment_id   string
	project_name     string
	project_id       string
	tag              string
	parameters       map[string]Parameter
	templates        map[string]Template
}

type Parameter struct {
	id     string
	name   string
	value  string
	secret bool
}

type Template struct {
	id    string
	name  string
	body  string
	value func() string
}

func NewCTApi(api_key string, api_url string, user_agent string, debug bool) *CTApi {
	log.Debug("Creating new CTApi")

	ctapi := new(CTApi)
	ctapi.configuration = cloudtruth.NewConfiguration()
	ctapi.configuration.UserAgent = user_agent
	u, err := url.Parse(api_url)
	if err != nil {
		log.Fatalf("Invalid url: %s, err: %v", api_url, err)
	}
	ctapi.configuration.Host = u.Host
	ctapi.configuration.Scheme = u.Scheme
	ctapi.configuration.Debug = debug

	ctapi.configuration.AddDefaultHeader("Authorization", fmt.Sprintf("Api-Key %s", api_key))

	ctapi.client = cloudtruth.NewAPIClient(ctapi.configuration)

	ctapi.environments = make(map[string]cloudtruth.Environment)
	ctapi.projects = make(map[string]cloudtruth.Project)

	return ctapi
}

func (ctapi *CTApi) Environments() map[string]cloudtruth.Environment {
	if len(ctapi.environments) == 0 {
		log.Debug("Fetching environments")

		resp, r, err := ctapi.client.EnvironmentsApi.EnvironmentsList(context.Background()).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.EnvironmentsList``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
			os.Exit(2)
		}

		for _, p := range resp.Results {
			ctapi.environments[p.Name] = p
		}
	}

	return ctapi.environments
}

func (ctapi *CTApi) EnvironmentId(name string) string {
	if len(ctapi.environments) != 0 {
		return ctapi.environments[name].Id
	}

	log.Debugf("Looking up environment id: %s", name)

	request := ctapi.client.EnvironmentsApi.EnvironmentsList(context.Background())
	request = request.Name(name)
	resp, r, err := request.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.EnvironmentsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		os.Exit(2)
	}

	if len(resp.Results) != 1 || resp.Results[0].Name != name {
		log.Errorf("Could not find environment: %s", name)
		os.Exit(2)
	}

	return resp.Results[0].Id
}

func (ctapi *CTApi) Projects() map[string]cloudtruth.Project {
	if len(ctapi.projects) == 0 {
		log.Debug("Fetching projects")

		resp, r, err := ctapi.client.ProjectsApi.ProjectsList(context.Background()).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsList``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
			os.Exit(2)
		}

		for _, p := range resp.Results {
			ctapi.projects[p.Name] = p
		}
	}

	return ctapi.projects
}

func (ctapi *CTApi) ProjectId(name string) string {
	if len(ctapi.projects) != 0 {
		return ctapi.projects[name].Id
	}

	log.Debugf("Looking up project id: %s", name)

	request := ctapi.client.ProjectsApi.ProjectsList(context.Background())
	request = request.Name(name)
	resp, r, err := request.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		os.Exit(2)
	}

	if len(resp.Results) != 1 || resp.Results[0].Name != name {
		log.Errorf("Could not find project: %s", name)
		os.Exit(2)
	}

	return resp.Results[0].Id
}

func NewCTProject(ctapi *CTApi, project_name string, environment_name string, tag string) *CTProject {
	ctproject := new(CTProject)
	ctproject.ctapi = ctapi
	ctproject.project_name = project_name
	ctproject.project_id = ctapi.ProjectId(project_name)
	ctproject.environment_name = environment_name
	ctproject.environment_id = ctapi.EnvironmentId(environment_name)
	ctproject.tag = tag
	ctproject.parameters = make(map[string]Parameter)
	ctproject.templates = make(map[string]Template)
	return ctproject
}

func (ctproject *CTProject) Parameters() map[string]Parameter {
	if len(ctproject.parameters) == 0 {
		log.Infof("Fetching parameters %s=%s/%s=%s",
			ctproject.project_name, ctproject.project_id,
			ctproject.environment_name, ctproject.environment_id)

		request := ctproject.ctapi.client.ProjectsApi.ProjectsParametersList(context.Background(), ctproject.project_id)
		request = request.Environment(ctproject.environment_id)
		request = request.Evaluate(true)
		if ctproject.tag != "" {
			request = request.Tag(ctproject.tag)
		}

		resp, r, err := request.Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsParametersList``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
			os.Exit(2)
		}

		for _, p := range resp.Results {
			v := ""
			if len(p.ValuesFlat) != 0 {
				v = *p.ValuesFlat[0].Value.Get()
			}
			ctproject.parameters[p.Name] = Parameter{p.Id, p.Name, v, *p.Secret}
		}
	}
	return ctproject.parameters
}

func (ctproject *CTProject) Template(template_name string) string {
	template := ctproject.templates[template_name]

	log.Infof("Fetching template %s=%s/%s=%s/%s=%s",
		ctproject.project_name, ctproject.project_id,
		ctproject.environment_name, ctproject.environment_id,
		template.name, template.id)

	request := ctproject.ctapi.client.ProjectsApi.ProjectsTemplatesRetrieve(context.Background(), template.id, ctproject.project_id)
	request = request.Environment(ctproject.environment_id)
	if ctproject.tag != "" {
		request = request.Tag(ctproject.tag)
	}

	resp, r, err := request.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsTemplatesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		os.Exit(2)
	}

	return *resp.Body
}

func (ctproject *CTProject) Templates() map[string]Template {
	if len(ctproject.templates) == 0 {
		log.Infof("Fetching templates %s=%s/%s=%s",
			ctproject.project_name, ctproject.project_id,
			ctproject.environment_name, ctproject.environment_id)

		request := ctproject.ctapi.client.ProjectsApi.ProjectsTemplatesList(context.Background(), ctproject.project_id)
		request = request.Environment(ctproject.environment_id)
		request = request.Evaluate(false)
		if ctproject.tag != "" {
			request = request.Tag(ctproject.tag)
		}

		resp, r, err := request.Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsTemplatesList``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
			os.Exit(2)
		}

		for _, t := range resp.Results {
			template_name := t.Name
			log.Debugf("Adding template reference '%s'", template_name)
			anon := func() string { return ctproject.Template(template_name) }
			ctproject.templates[t.Name] = Template{t.Id, t.Name, "", anon}
		}
	}

	return ctproject.templates
}

func (template *Template) Get() string {
	if template.body == "" {
		template.body = template.value()
	}
	return template.body
}
