dist/argocd-cloudtruth-plugin: pkg/cloudtruth/client.go *.go
	go build -o dist/argocd-cloudtruth-plugin

client: pkg/cloudtruth/client.go

pkg/cloudtruth/client.go: pkg/openapi.yml
	docker run --rm \
		-v "$(shell pwd)/pkg:/pkg" \
		--user "$(shell id -u):$(shell id -g)" \
		openapitools/openapi-generator-cli:v6.2.0 generate \
		-i /pkg/openapi.yml \
		-g go \
		-o /pkg/cloudtruth \
		--additional-properties packageName=cloudtruth \
		--additional-properties isGoSubmodule=true \
		--additional-properties packageVersion=1.0.0 \
		--additional-properties enumClassPrefix=true \
		--type-mappings=object=interface{}

pkg/openapi.yml: pkg
	curl -s https://api.cloudtruth.io/api/schema/ > pkg/openapi.yml

pkg:
	mkdir -p pkg

clean:
	rm -rf pkg dist
