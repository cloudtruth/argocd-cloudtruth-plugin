module github.com/argocd-cloudtruth-plugin

go 1.17

require (
	github.com/argocd-cloudtruth-plugin/pkg/cloudtruth v1.0.0
	github.com/jessevdk/go-flags v1.5.0
	github.com/sirupsen/logrus v1.9.2
	github.com/yargevad/filepathx v1.0.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/oauth2 v0.8.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)

replace github.com/argocd-cloudtruth-plugin/pkg/cloudtruth v1.0.0 => ./pkg/cloudtruth
