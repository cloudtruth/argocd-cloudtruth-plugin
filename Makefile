pkg/openapi.yml:
	mkdir -p pkg
	curl https://api.cloudtruth.io/api/schema/ > pkg/openapi.yml

client: pkg/openapi.yml
	docker run --rm \
		-v "$(shell pwd)/pkg:/pkg" \
		--user "$(shell id -u):$(shell id -g)" \
		-e GO_POST_PROCESS_FILE="gofmt -s -w" \
		openapitools/openapi-generator-cli generate \
		-i /pkg/openapi.yml \
		-g go \
		-o /pkg/cloudtruth \
		--additional-properties packageName=cloudtruth \
		--additional-properties isGoSubmodule=true \
		--additional-properties packageVersion=1.0.0 \
		--additional-properties enumClassPrefix=true \
		--type-mappings=object=interface{} \
    	--enable-post-process-file \

clean:
	rm -rf pkg
