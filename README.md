# argocd-cloudtruth-plugin

The CloudTruth configuration management plugin for [ArgoCD](https://argo-cd.readthedocs.io/en/stable/)  When attached to an argocd project, it replaces parameter and secret references (`<parameter_name>`) with the values looked up from CloudTruth.

## Installation

ArgoCD plugin installation is somewhat of a manual process, you can either use the `install/*.yaml` files as a guide, or run `install.sh` to use kubectl patch to apply them:

```shell
git clone https://github.com/cloudtruth/argocd-cloudtruth-plugin
cd argocd-cloudtruth-plugin/install
./install.sh
```

## Usage


| Parameter | Description | Type | Default | Required |
|-----------|-------------|------|---------|:--------:|
|  |  | string | n/a | yes |

## Development

After checking out the repo, run `make client` to generate the go client stubs for the cloudtruth rest api.

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/cloudtruth/argocd-cloudtruth-plugin
