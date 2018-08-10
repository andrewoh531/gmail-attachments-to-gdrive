# Variables
cfn_output_file=cfn-template-output.yaml
app_name=gmail-attachments-to-gdrive
artifact_name=gmail-attacher

.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./$(artifact_name)
	
build:
	GOOS=linux GOARCH=amd64 go build -o src/$(artifact_name) ./src/

test:
	go test -v ./src/

run:
	sam local start-api

package:
	sam package \
		--template-file template.yaml \
		--output-template-file $(cfn_output_file) \
		--s3-bucket $(app_name)

deploy:
	sam deploy \
		--template-file $(cfn_output_file) \
		--stack-name $(app_name) \
		--capabilities CAPABILITY_IAM \
		--parameter-overrides GmailSearchString=yolobaby \
		                      GoogleDriveUploadFolder=gdrive \
		                      GmailOAuthToken=andrew1 \
		                      GoogleDriveOAuthToken=password

invoke:
	aws lambda invoke --function-name gmail-attachments-to-gdri-RetrieveFromGmailUploadT-17R7DXPILYMG4 output.txt
