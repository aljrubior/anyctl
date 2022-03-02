package clients

import (
	"encoding/json"
	"fmt"
	error2 "github.com/aljrubior/anyctl/errors"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

const (
	CONTENT_TYPE_APPLICATION_JSON_VALUE    = "application/json"
	CONTENT_TYPE_MULTIPART_FORM_DATA_VALUE = "multipart/form-data"
)

type HttpClient struct {
}

func (_ *HttpClient) GetBearerTokenValue(token string) string {
	return fmt.Sprintf("%s %s", "bearer", token)
}

func (_ *HttpClient) ThrowError(response *http.Response) error {

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		errors.New(response.Status)
	}

	var body map[string]string

	if err := json.Unmarshal(data, &body); err != nil {
		errors.New(response.Status)
	}

	switch response.StatusCode {
	case 401:
		return error2.NewUnauthorizedError(response)
	default:
		return errors.New(fmt.Sprintf("%s. Reason: %s", response.Status, string(data)))
	}
}

func (_ *HttpClient) ThrowErrorWithReason(response *http.Response, data []byte) error {

	switch response.StatusCode {
	case 401:
		return error2.NewUnauthorizedError(response)
	default:
		return errors.New(response.Status)
	}
}

func (this *HttpClient) AddDefaultHeaders(req *http.Request, orgId, envId, token string) {

	req.Header.Add("Authorization", this.GetBearerTokenValue(token))
	req.Header.Add("X-ANYPNT-ENV-ID", envId)
	req.Header.Add("X-ANYPNT-ORG-ID", orgId)
}

func (this *HttpClient) AddAuthorizationHeader(req *http.Request, token string) {

	req.Header.Add("Authorization", this.GetBearerTokenValue(token))
}

func (this *HttpClient) AddContentTypeApplicationJsonHeader(req *http.Request) {

	this.AddContentTypeHeader(req, "application/json")
}

func (this *HttpClient) AddContentTypeHeader(req *http.Request, contentType string) {

	req.Header.Add("Content-Type", contentType)
}

func (this *HttpClient) AddCustomHeader(req *http.Request, withKey, withValue string) {

	req.Header.Add(withKey, withValue)
}
