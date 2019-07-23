package handlers

import (
	"context"

	log "github.com/sirupsen/logrus"

	"net/http"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var (
	// you need to generate personal access token at
	// https://github.com/settings/applications#personal-access-tokens
	personalAccessToken = os.Getenv("TOKEN")
)

type TokenSource struct {
	AccessToken string
}

type config struct {
	URL     string `json:"url"`
	Content string `json:"content_type"`
	Secret  string `json:"secret"`
	NoSsl   string `json:"insecure_ssl"`
}

type payload struct {
	Name   string   `json:"name"`
	Config config   `json:"config"`
	Events []string `json:"events"`
	Active bool     `json:"active"`
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

type Creator interface {
	NewClient(httpClient *http.Client) *github.Client
}

type WebHookCreator struct {
}

func (wh WebHookCreator) CreateWebhook(w http.ResponseWriter, req *http.Request) {
	tokenSource := &TokenSource{
		AccessToken: personalAccessToken,
	}
	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	client := github.NewClient(oauthClient)
	config := config{
		URL:     "http://httpbin.org/post/webhook",
		Content: "json",
		Secret:  "my-secret-key",
		NoSsl:   "1"}
	payload := payload{
		Name:   "web",
		Config: config,
		Events: []string{"star", "push"},
		Active: true,
	}
	req, err := client.NewRequest(http.MethodPost, "/repos/kyma-incubator/test-k8s/hooks", payload)
	if err != nil {
		log.Info(err)
	}

	resp, err := client.Do(context.Background(), req, nil)
	if err != nil {
		log.Print(err)
	}
	log.Println(resp)
}
