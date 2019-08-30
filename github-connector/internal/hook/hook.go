package hook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/apperrors"
)

const (
	kymaURLPrefix = "https://"
	kymaURLSuffix = "/webhook"
	kymaURLFormat = "%s%s%s"
	charset       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

//Hook is a struct that contains information about github's repo/org url, OAuth token and allows creating webhooks
type Hook struct {
	kymaURL string
}

//NewHook create Hook structure
func NewHook(URL string) Hook {
	kURL := fmt.Sprintf(kymaURLFormat, kymaURLPrefix, URL, kymaURLSuffix)
	return Hook{kymaURL: kURL}
}

//Create build request and create webhook in github's repository or organization
func (c Hook) Create(t string, githubURL string) (string, apperrors.AppError) {
	token := fmt.Sprintf("token %s", t)
	secret := createSecret(charset)
	hook := PayloadDetails{
		Name:   "web",
		Active: true,
		Config: Config{
			URL:         c.kymaURL,
			InsecureSSL: "1",
			ContentType: "json",
			Secret:      secret,
		},
		Events: []string{"*"},
	}

	payloadJSON, err := json.Marshal(hook)
	if err != nil {
		return "", apperrors.Internal("Failed to marshal hook: %s", err.Error())
	}

	requestReader := bytes.NewReader(payloadJSON)
	httpRequest, err := http.NewRequest(http.MethodPost, githubURL, requestReader)

	if err != nil {
		return "", apperrors.Internal("Failed to create JSON request: %s", err.Error())
	}

	httpRequest.Header.Set("Authorization", token)

	client := &http.Client{}
	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		return "", apperrors.UpstreamServerCallFailed("Failed to make request to '%s': %s", githubURL, err.Error())
	}

	if httpResponse.StatusCode != http.StatusCreated {
		return "", apperrors.UpstreamServerCallFailed("Unpredicted response code: %v", httpResponse.StatusCode)
	}
	return secret, nil
}

func createSecret(charset string) string {
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	secret := make([]byte, (rand.Intn(7) + 8))
	for i := range secret {
		secret[i] = charset[seed.Intn(len(charset))]
	}
	return string(secret)
}
