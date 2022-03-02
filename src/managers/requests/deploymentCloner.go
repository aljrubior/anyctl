package requests

import (
	"encoding/json"
	"github.com/aljrubior/anyctl/clients/deployments/response"
	"strings"
)

func NewDeploymentRequestBuilder(fromResponse *response.DeploymentResponse, withName string) *DeploymentRequestBuilder {
	return &DeploymentRequestBuilder{
		fromResponse,
		withName,
	}
}

type DeploymentRequestBuilder struct {
	response *response.DeploymentResponse
	withName string
}

func (this *DeploymentRequestBuilder) Clone() (*DeploymentRequest, error) {

	return this.cloneResponse()
}

func (this *DeploymentRequestBuilder) cloneResponse() (*DeploymentRequest, error) {

	currentName := this.response.Name

	data, err := json.Marshal(this.response)

	if err != nil {
		return nil, err
	}

	responseAsString := string(data)

	responseAfterReplace := strings.ReplaceAll(responseAsString, currentName, this.withName)

	var clone DeploymentRequest

	if err := json.Unmarshal([]byte(responseAfterReplace), clone); err != nil {
		return nil, err
	}

	return &clone, nil
}
