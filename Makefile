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
		--parameter-overrides \
            GmailSearchString=yolobaby \
            GmailClientCredentialName=gmail-client-credentials \
            GoogleDriveUploadFolder=gdrive \
            GmailOAuthRefreshTokenName=tax-helper-gmail-refresh-token-name \
            GoogleDriveOAuthToken=password

invoke:
	aws lambda invoke --function-name gmail-attachments-to-gdri-RetrieveFromGmailUploadT-17R7DXPILYMG4 output.log && cat output.log

invoke-local:
	echo '{}' | sam local invoke --env-vars env.json "RetrieveFromGmailUploadToGDrive" --log-file output.log && cat output.log

.PHONY: deps clean build package deploy invoke invoke-local