package hooks

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/apperrors"
)

const (
	linkPrefix = "http://api.github.com/"
	linkSuffix = "/hooks"
	linkFormat = "%s%s%s"
)

type creator struct {
	token   string
	repoURL string
}

func NewCreator(t string, rURL string) creator {
	return creator{token: t, repoURL: rURL}
}

func (c creator) Create() error {

	payloadJSON := []byte(`{
		"name": "web",
		"active": true,
		"config": {
		  "url": "http://example1234.com/webhook"
		}
	  }`)

	url := fmt.Sprintf(linkFormat, linkPrefix, c.repoURL, linkSuffix)

	requestReader := bytes.NewReader(payloadJSON)

	httpRequest, err := http.NewRequest(http.MethodPost, url, requestReader)

	if err != nil {
		return apperrors.Internal("Failed to create JSON request: %s", err.Error())
	}

	httpRequest.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	httpResponse, err := client.Do(httpRequest)

	if err != nil {
		return apperrors.UpstreamServerCallFailed("Failed to make request to '%s': %s", url, err.Error())
	}

	if httpResponse.StatusCode != http.StatusOK {
		return apperrors.UpstreamServerCallFailed("Incorrect response code '%d' while sending JSON request from %s", httpResponse.StatusCode, url)
	}

	return nil
}
