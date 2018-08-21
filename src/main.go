package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/andrewoh531/gmail-attachments-to-gdrive/src/clients"
)

// For mocking
var retrieveParameter = clients.RetrieveFromParameterStore


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
	ssmClient := clients.GetSsmClient()
	gmailOAuthToken := retrieveParameter(ssmClient, os.Getenv("GMAIL_OAUTH_TOKEN"))
	googleDriveOAuthToken := retrieveParameter(ssmClient, os.Getenv("GOOGLE_DRIVE_OAUTH_TOKEN"))

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
