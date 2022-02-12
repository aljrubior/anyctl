package requests

import (
	"bytes"
	"fmt"
	"github.com/aljrubior/anyctl/clients"
	"github.com/aljrubior/anyctl/conf"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func NewPostUploadAssetRequest(
	config conf.AssetClientConfig,
	bearerToken,
	organizationId,
	assetName,
	assetVersion,
	fromFilePath string) *PostUploadAssetRequest {
	return &PostUploadAssetRequest{
		config:         config,
		bearerToken:    bearerToken,
		organizationId: organizationId,
		assetName:      assetName,
		assetVersion:   assetVersion,
		fromFilePath:   fromFilePath,
	}
}

type PostUploadAssetRequest struct {
	clients.BaseHttpRequest
	config         conf.AssetClientConfig
	bearerToken    string
	organizationId string
	assetName      string
	assetVersion   string
	fromFilePath   string
}

func (this *PostUploadAssetRequest) buildUri() string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := fmt.Sprintf(this.config.UploadAssetPath, this.organizationId, this.organizationId, this.assetName, this.assetVersion)

	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *PostUploadAssetRequest) Build() (*http.Request, error) {

	uri := this.buildUri()

	body, contentType, err := this.buildMultiPartBody()

	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest(http.MethodPost, uri, body)

	this.AddAuthorizationHeader(req, this.bearerToken)
	this.AddContentType(req, *contentType)

	return req, nil
}

func (this *PostUploadAssetRequest) buildMultiPartBody() (*bytes.Buffer, *string, error) {

	file, err := os.Open(this.fromFilePath)

	defer file.Close()

	if err != nil {
		return nil, nil, err
	}

	if err != nil {
		return nil, nil, err
	}

	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("files.mule-application.jar", fmt.Sprintf("@%s", this.fromFilePath))

	if err != nil {
		return nil, nil, err
	}

	io.Copy(part, file)

	writer.WriteField("name", this.assetName)

	err = writer.Close()

	if err != nil {
		return nil, nil, err
	}

	contentType := writer.FormDataContentType()

	return body, &contentType, nil
}
