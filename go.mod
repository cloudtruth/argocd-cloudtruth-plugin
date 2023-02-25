module github.com/argocd-cloudtruth-plugin

go 1.17

require (
	github.com/argocd-cloudtruth-plugin/pkg/cloudtruth v1.0.0
	github.com/jessevdk/go-flags v1.5.0
	github.com/sirupsen/logrus v1.8.1
	github.com/yargevad/filepathx v1.0.0
)

require (
	github.com/golang/protobuf v1.4.2 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8 // indirect
	golang.org/x/sys v0.5.0 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
)

replace github.com/argocd-cloudtruth-plugin/pkg/cloudtruth v1.0.0 => ./pkg/cloudtruth
