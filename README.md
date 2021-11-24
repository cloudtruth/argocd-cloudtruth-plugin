# argocd-cloudtruth-plugin

The CloudTruth configuration management plugin for [ArgoCD](https://argo-cd.readthedocs.io/en/stable/)  When attached to an argocd project, it replaces parameter and secret references (`<parameter_name>`) with the values looked up from CloudTruth.

## Installation

ArgoCD plugin installation is somewhat of a manual process, you can either use the `install/*.yaml` files as a guide, or run `install/setup.sh` to use kubectl patch to apply them:

```shell
curl -s https://raw.githubusercontent.com/cloudtruth/argocd-cloudtruth-plugin/main/install/setup.sh | sh
```
OR
```shell
git clone https://github.com/cloudtruth/argocd-cloudtruth-plugin
./argocd-cloudtruth-plugin/install/setup.sh
```

## Usage

The default installation process adds the following settings to the Secret named argocd-cloudtruth-plugin, and applies them as environment variables to the argocd-repo-server Deployment.  Theses act as defaults for the plugin, and can be overriden for each argocd project in the plugin setup page for each project.

| Parameter | Description | Type | Default | Required |
|-----------|-------------|------|---------|:--------:|
| CLOUDTRUTH_API_KEY | the api key for authenticating to the CloudTruth Rest API | string | n/a | yes |
| CLOUDTRUTH_ENVIRONMENT | the CloudTruth environment to query against | string | `default` | no |
| CLOUDTRUTH_PROJECT | the CloudTruth project to query against | string | `MyFirstProject` | yes |
| CLOUDTRUTH_TAG | the CloudTruth tag to query against | string | `""` | no |
| CLOUDTRUTH_REFERENCE_PATTERN | the pattern that indicates parameter references in the input files | string | `<%s>` | no |
| CLOUDTRUTH_FILE_PATTERN | the file pattern (glob) of the input files | array(string) delim=, | `*.y*ml` | no |

## Development

After checking out the repo, run `make`.  To generate new client stubs (vs what is checked in), run `make clean` before running `make`

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/cloudtruth/argocd-cloudtruth-plugin
