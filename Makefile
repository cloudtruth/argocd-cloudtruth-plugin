build/bin/argocd-cloudtruth-plugin: build/cloudtruth
	go build -o build/bin/argocd-cloudtruth-plugin

client: build/cloudtruth

build/cloudtruth: build/openapi.yml
	docker run --rm \
		-v "$(shell pwd)/build:/build" \
		--user "$(shell id -u):$(shell id -g)" \
		openapitools/openapi-generator-cli generate \
		-i /build/openapi.yml \
		-g go \
		-o /build/cloudtruth \
		--additional-properties packageName=cloudtruth \
		--additional-properties isGoSubmodule=true \
		--additional-properties packageVersion=1.0.0 \
		--additional-properties enumClassPrefix=true \
		--type-mappings=object=interface{}

build/openapi.yml: build
	wget -O build/openapi.yml https://api.cloudtruth.io/api/schema/

build:
	mkdir -p build

clean:
	rm -rf pkg bin
