package clients

import (
	"github.com/aws/aws-sdk-go/service/ssm"
	"sync"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
)

var ssmClient *ssm.SSM
var once sync.Once

// Retrieve a single instance of an SSM Client
func GetSsmClient() *ssm.SSM {
	once.Do(func() {
		sess := session.Must(session.NewSession())
		ssmClient = ssm.New(sess, aws.NewConfig())
	})
	return ssmClient
}
