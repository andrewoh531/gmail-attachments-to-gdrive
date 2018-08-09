package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
)

var (
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

	// ErrNoIP No IP found in response
	ErrNoIP = errors.New("No IP in HTTP response")

	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	resp, err := http.Get(DefaultHTTPGetAddress)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if resp.StatusCode != 200 {
		return events.APIGatewayProxyResponse{}, ErrNon200Response
	}

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if len(ip) == 0 {
		return events.APIGatewayProxyResponse{}, ErrNoIP
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("GMAIL_OAUTH_TOKEN=%v, GMAIL_SEARCH_QUERY=%v, GOOGLE_DRIVE_OAUTH_TOKEN=%v, GOOGLE_DRIVE_UPLOAD_FOLDER=%v",
			os.Getenv("GMAIL_OAUTH_TOKEN"),
			os.Getenv("GMAIL_SEARCH_QUERY"),
			os.Getenv("GOOGLE_DRIVE_OAUTH_TOKEN"),
			os.Getenv("GOOGLE_DRIVE_UPLOAD_FOLDER")),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
