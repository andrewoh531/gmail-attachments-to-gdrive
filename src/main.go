package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"

	"personal/gmail-attachments-to-gdrive/src/clients"
)

func retrieveFromParameterStore(key string) string {

	output, err := clients.GetSsmClient().GetParameter(&ssm.GetParameterInput{
		Name:           &key,
		WithDecryption: aws.Bool(true),
	})

	if err != nil {
		fmt.Printf("Error retrieving SSM (%s): %v", key, err)
		os.Exit(1)
	}

	return *output.Parameter.Value
}

func handler(request events.CloudWatchEvent) (events.APIGatewayProxyResponse, error) {

	/*
		1 - Retrieve Gmail and Google Drive tokens from AWS Parameter Store
		2 - Query Gmail:
			a - Use specific date time range
			b - for any attachments using the provided gmail search string
		3 - For each match:
			a - Connect to Google Drive and upload using date of email
			b - Save information in a map/list
		4 - Send confirmation email for the google drive account of files uploaded
	 */

	gmailOAuthToken := retrieveFromParameterStore(os.Getenv("GMAIL_OAUTH_TOKEN"))
	googleDriveOAuthToken := retrieveFromParameterStore(os.Getenv("GOOGLE_DRIVE_OAUTH_TOKEN"))

	return events.APIGatewayProxyResponse{
		Body: fmt.Sprintf("GMAIL_OAUTH_TOKEN=%v, GMAIL_SEARCH_QUERY=%v, GOOGLE_DRIVE_OAUTH_TOKEN=%v, GOOGLE_DRIVE_UPLOAD_FOLDER=%v",
			gmailOAuthToken,
			os.Getenv("GMAIL_SEARCH_QUERY"),
			googleDriveOAuthToken,
			os.Getenv("GOOGLE_DRIVE_UPLOAD_FOLDER")),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
