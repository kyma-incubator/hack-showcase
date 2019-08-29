package hooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/apperrors"
	log "github.com/sirupsen/logrus"
)

const (
	kymaURLPrefix = "https://"
	kymaURLSuffix = "/webhook"
	kymaURLFormat = "%s%s%s"
)

//Creator is an struct that contain informations about github's repo/org url, OAuth token and allow creating webhooks
type Creator struct {
	token   string
	repoURL string
}

//NewCreator create Creator structure
func NewCreator(t string, rURL string) Creator {
	return Creator{token: t, repoURL: rURL}
}

//Create build request and create webhook in github's repository or organization
func (c Creator) Create(kURL string) apperrors.AppError {
	githubURL := c.repoURL
	kymaURL := fmt.Sprintf(kymaURLFormat, kymaURLPrefix, kURL, kymaURLSuffix)
	token := "token " + c.token
	hook := HookJSON{
		Name:   "web",
		Active: true,
		Config: &Config{
			URL:         kymaURL,
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

	if httpResponse.StatusCode != http.StatusCreated {
		return apperrors.UpstreamServerCallFailed("Failed to make request to '%s': %s", githubURL, err.Error())
	}

	if err != nil {
		return apperrors.UpstreamServerCallFailed("Failed to make request to '%s': %s", githubURL, err.Error())
	}

	log.Info("Webhook created!")

	return nil
}
