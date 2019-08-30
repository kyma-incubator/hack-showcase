package hook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/apperrors"
)

const (
	kymaURLPrefix = "https://"
	kymaURLSuffix = "/webhook"
	kymaURLFormat = "%s%s%s"
)

//Hook is an struct that contain informations about github's repo/org url, OAuth token and allow creating webhooks
type Hook struct {
	kymaURL string
}

//NewHook create Hook structure
func NewHook(URL string) Hook {
	kURL := fmt.Sprintf(kymaURLFormat, kymaURLPrefix, URL, kymaURLSuffix)
	return Hook{kymaURL: kURL}
}

//Create build request and create webhook in github's repository or organization
func (c Hook) Create(t string, githubURL string) apperrors.AppError {
	token := "token " + t
	hook := HookDetails{
		Name:   "web",
		Active: true,
		Config: &Config{
			URL:         c.kymaURL,
			InsecureSSL: "1",
			ContentType: "json",
		},
		Events: []string{"*"},
	}

	payloadJSON, err := json.Marshal(hook)
	if err != nil {
		return apperrors.Internal("Failed to marshal hook: %s", err.Error())
	}

	requestReader := bytes.NewReader(payloadJSON)
	httpRequest, err := http.NewRequest(http.MethodPost, githubURL, requestReader)

	if err != nil {
		return apperrors.Internal("Failed to create JSON request: %s", err.Error())
	}

	httpRequest.Header.Set("Authorization", token)

	client := &http.Client{}
	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		return apperrors.UpstreamServerCallFailed("Failed to make request to '%s': %s", githubURL, err.Error())
	}

	if httpResponse.StatusCode != http.StatusCreated {
		return apperrors.UpstreamServerCallFailed("Bad response code. Maybe webhook already exist?")
	}
	return nil
}
