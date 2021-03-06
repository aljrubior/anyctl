package targets

import (
	"encoding/json"
	"fmt"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/clients/targets/requests"
	"github.com/aljrubior/anyctl/clients/targets/response"
	"github.com/aljrubior/anyctl/conf"
	"io/ioutil"
	"net/http"
	"time"
)

func NewDefaultTargetClient(config conf.TargetClientConfig) DefaultTargetClient {
	return DefaultTargetClient{
		config: config,
	}
}

type DefaultTargetClient struct {
	clients.HttpClient
	config conf.TargetClientConfig
}

func (this DefaultTargetClient) GetTargets(orgId, envId, token string) (*response.TargetsResponse, error) {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetTargets(&this.config, orgId, envId, token).Build()

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

	return this.buildTargetsResponse(data)
}

func (this DefaultTargetClient) buildTargetsResponse(data []byte) (*response.TargetsResponse, error) {

	var targets map[string][]map[string]interface{}

	err := json.Unmarshal(data, &targets)

	if err != nil {
		return nil, err
	}

	var targetsResponse []response.TargetResponse
	var targetResponse response.TargetResponse

	for _, v := range targets["data"] {

		kind := fmt.Sprintf("%v", v["type"])

		switch kind {
		case "MC":

			runtimeFabricTargetResponse := new(response.RuntimeFabricTargetResponse)

			itemAsBytes, err := json.Marshal(v)

			if err != nil {
				return nil, err
			}

			json.Unmarshal(itemAsBytes, &runtimeFabricTargetResponse)

			targetResponse = runtimeFabricTargetResponse

			targetsResponse = append(targetsResponse, targetResponse)

		case "SERVER":

			standaloneTargetResponse := new(response.StandaloneTargetResponse)

			itemAsBytes, err := json.Marshal(v)

			if err != nil {
				return nil, err
			}

			json.Unmarshal(itemAsBytes, &standaloneTargetResponse)

			targetResponse = standaloneTargetResponse

			targetsResponse = append(targetsResponse, targetResponse)
		}
	}

	return &response.TargetsResponse{
		Data: targetsResponse,
	}, nil
}
