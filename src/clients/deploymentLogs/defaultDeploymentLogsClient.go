package deploymentLogs

import (
	"encoding/json"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/clients/deploymentLogs/requests"
	"github.com/aljrubior/anyctl/clients/deploymentLogs/response"
	"github.com/aljrubior/anyctl/conf"
	"io/ioutil"
	"net/http"
	"time"
)

func NewDefaultDeploymentLogsClient(config conf.DeploymentLogsClientConfig) DefaultDeploymentLogsClient {

	return DefaultDeploymentLogsClient{
		config: config,
	}

}

type DefaultDeploymentLogsClient struct {
	clients.HttpClient
	config conf.DeploymentLogsClientConfig
}

func (this DefaultDeploymentLogsClient) GetLogs(orgId, envId, token, deploymentId, specId string) (*[]response.DeploymentLogMessageResponse, error) {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetLogsRequest(&this.config, token, orgId, envId, deploymentId, specId).Build()

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, this.ThrowError(resp)
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var response []response.DeploymentLogMessageResponse

	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
