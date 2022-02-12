package assets

import (
	"encoding/json"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/clients/assets/requests"
	"github.com/aljrubior/anyctl/clients/assets/response"
	"github.com/aljrubior/anyctl/conf"
	"io/ioutil"
	"net/http"
	"time"
)

func NewDefaultAssetClient(config conf.AssetClientConfig) *DefaultAssetClient {
	return &DefaultAssetClient{
		config: config,
	}
}

type DefaultAssetClient struct {
	clients.HttpClient
	AssetClient
	config conf.AssetClientConfig
}

func (this *DefaultAssetClient) FindAssets(orgId, envId, token, assetName string) (*[]response.AssetResponse, error) {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetAssetsRequest(this.config, token, orgId, envId, assetName).Build()

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, this.ThrowError(resp)
	}

	var assets []response.AssetResponse

	err = json.Unmarshal(data, &assets)

	if err != nil {
		return nil, err
	}

	return &assets, nil
}

func (this *DefaultAssetClient) FindLatestVersion(token, groupId, assetName string) (*response.AssetResponse, error) {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetAssetLatestVersionRequest(this.config, token, groupId, assetName).Build()

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, this.ThrowError(resp)
	}

	var asset response.AssetResponse

	err = json.Unmarshal(data, &asset)

	if err != nil {
		return nil, err
	}

	return &asset, nil
}

func (this *DefaultAssetClient) FindSpecificVersion(token, groupId, assetName, version string) (*response.AssetResponse, error) {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetAssetVersionRequest(this.config, token, groupId, assetName, version).Build()

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

	var asset response.AssetResponse

	err = json.Unmarshal(data, &asset)

	if err != nil {
		return nil, err
	}

	return &asset, nil
}

func (this *DefaultAssetClient) UploadArtifact(token, orgId, filePath, assetName, version string) (*response.AssetPublicationResponse, error) {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req, err := requests.NewPostUploadAssetRequest(this.config, token, orgId, assetName, version, filePath).Build()

	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 202 {
		return nil, this.ThrowError(resp)
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var asset response.AssetPublicationResponse

	err = json.Unmarshal(data, &asset)

	if err != nil {
		return nil, err
	}

	return &asset, nil
}
