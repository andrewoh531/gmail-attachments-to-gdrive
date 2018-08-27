# Variables
cfn_output_file=cfn-template-output.yaml
app_name=gmail-attachments-to-gdrive
artifact_name=gmail-attacher

deps:
	go get -u ./...

clean: 
	rm -rf ./$(artifact_name)
	
build:
	GOOS=linux GOARCH=amd64 go build -o src/$(artifact_name) ./src/

test:
	go test -v ./src/...

run:
	sam local start-api

package: build
	sam package \
		--template-file template.yaml \
		--output-template-file $(cfn_output_file) \
		--s3-bucket $(app_name)

deploy: package
	sam deploy \
		--template-file $(cfn_output_file) \
		--stack-name $(app_name) \
		--capabilities CAPABILITY_IAM \
		--parameter-overrides GmailSearchString=yolobaby \
		                      GoogleDriveUploadFolder=gdrive \
		                      GmailOAuthAccessTokenName=tax-helper-gmail-access-token-name \
		                      GmailOAuthRefreshTokenName=tax-helper-gmail-refresh-token-name \
		                      GoogleDriveOAuthToken=password

invoke:
	aws lambda invoke --function-name gmail-attachments-to-gdri-RetrieveFromGmailUploadT-17R7DXPILYMG4 output.txt

invoke-local:
	echo '{"message": "Hey, are you there?" }' | sam local invoke --env-vars env.json "RetrieveFromGmailUploadToGDrive"


.PHONY: deps clean build invoke invoke-local