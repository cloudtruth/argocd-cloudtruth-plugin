module github.com/argocd-cloudtruth-plugin

go 1.17

require github.com/argocd-cloudtruth-plugin/pkg/cloudtruth v1.0.0

require (
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/jessevdk/go-flags v1.5.0 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/yargevad/filepathx v1.0.0 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8 // indirect
	golang.org/x/sys v0.0.0-20210320140829-1e4c9ba3b0c4 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
)

replace github.com/argocd-cloudtruth-plugin/pkg/cloudtruth v1.0.0 => ./build/cloudtruth
