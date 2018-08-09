# Variables
cfn_output_file=cfn-template-output.yaml

.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./hello-world/hello-world
	
build:
	GOOS=linux GOARCH=amd64 go build -o hello-world/hello-world ./hello-world

package:
	sam package \
		--template-file template.yaml \
		--output-template-file $(cfn_output_file) \
		--s3-bucket tax-file-aggregator

deploy:
	sam deploy \
		--template-file $(cfn_output_file) \
		--stack-name tax-file-aggregator \
		--capabilities CAPABILITY_IAM