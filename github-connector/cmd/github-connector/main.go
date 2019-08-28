package main

import (
	"net/http"
	"os"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/github"
	"github.com/kyma-incubator/hack-showcase/github-connector/internal/hooks"
	"github.com/kyma-incubator/hack-showcase/github-connector/internal/registration"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/events"
	"github.com/kyma-incubator/hack-showcase/github-connector/internal/handlers"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("server started")

	builder := registration.NewPayloadBuilder(registration.NewFileReader(), os.Getenv("GITHUB_CONNECTOR_NAME"))
	id, err := registration.NewApplicationRegistryClient(builder, 5, 10).RegisterService()

	if err != nil {
		log.Fatal("Fatal error: ", err.Error())
	}
	log.WithFields(log.Fields{
		"id": id,
	}).Info("Service registered")

	creator := hooks.NewCreator(os.Getenv("GITHUB_TOKEN"), os.Getenv("GITHUB_REPO_URL"))
	creator.Create()

	kyma := events.NewSender(&http.Client{}, events.NewValidator(), "http://event-publish-service.kyma-system:8080/v1/events")
	webhook := handlers.NewWebHookHandler(
		github.ReceivingEventsWrapper{},
		kyma,
	)

	http.HandleFunc("/webhook", webhook.HandleWebhook)
	log.Info(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
