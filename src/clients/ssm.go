package clients

import (
	"github.com/aws/aws-sdk-go/service/ssm"
	"sync"
	"github.com/aws/aws-sdk-go/aws/session"
	"fmt"
	"os"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
	"github.com/aws/aws-sdk-go/aws"
)

var ssmClient *ssm.SSM
var once sync.Once

// Retrieve a single instance of an SSM Client
func GetSsmClient() *ssm.SSM {
	once.Do(func() {
		sess := session.Must(session.NewSession())
		ssmClient = ssm.New(sess)
	})
	return ssmClient
}

func RetrieveFromParameterStore(ssmClient ssmiface.SSMAPI, key string) string {

	output, err := ssmClient.GetParameter(&ssm.GetParameterInput{
		Name:           &key,
		WithDecryption: aws.Bool(true),
	})

	if err != nil {
		fmt.Printf("Error retrieving SSM (%s): %v", key, err)
		os.Exit(1)
	}

	return *output.Parameter.Value
}
