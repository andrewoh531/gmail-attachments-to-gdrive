package clients

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/pkg/errors"
	)

var expectedGetParamResponse = "Sample-Value"

type mockSSMClient struct {
	ssmiface.SSMAPI
}

func (m *mockSSMClient) GetParameter(param *ssm.GetParameterInput) (*ssm.GetParameterOutput, error) {
	if *param.Name == "err" {
		return nil, errors.New("Error")
	}

	output := ssm.GetParameterOutput{}
	output.SetParameter(&ssm.Parameter{Value: &expectedGetParamResponse})
	return &output, nil
}

func TestHandler(t *testing.T) {

	//t.Run("When retrieval fails then it should log and exit", func(t *testing.T) {
	//
	//	response := RetrieveFromParameterStore(&mockSSMClient{}, "err")
	//
	//})

	t.Run("Should successfully return parameter value", func(t *testing.T) {

		response := RetrieveFromParameterStore(&mockSSMClient{}, "sample-key")

		if response != expectedGetParamResponse {
			t.Errorf("Expected response of %s but was %s.", expectedGetParamResponse, response)
		}
	})
}
