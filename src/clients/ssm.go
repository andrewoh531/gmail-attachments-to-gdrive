package clients

import (
	"github.com/aws/aws-sdk-go/service/ssm"
	"sync"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"fmt"
	"os"
)

var ssmClient *ssm.SSM
var once sync.Once

// Retrieve a single instance of an SSM Client
func getSsmClient() *ssm.SSM {
	once.Do(func() {
		sess := session.Must(session.NewSession())
		ssmClient = ssm.New(sess, aws.NewConfig())
	})
	return ssmClient
}

func RetrieveFromParameterStore(key string) string {

	output, err := getSsmClient().GetParameter(&ssm.GetParameterInput{
		Name:           &key,
		WithDecryption: aws.Bool(true),
	})

	if err != nil {
		fmt.Printf("Error retrieving SSM (%s): %v", key, err)
		os.Exit(1)
	}

	return *output.Parameter.Value
}
