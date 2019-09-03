package main

import (
	"net/http"
	"os"

	"github.com/kyma-incubator/hack-showcase/slack-connector/internal/events"
	"github.com/kyma-incubator/hack-showcase/slack-connector/internal/handlers"
	"github.com/kyma-incubator/hack-showcase/slack-connector/internal/registration"
	"github.com/kyma-incubator/hack-showcase/slack-connector/internal/slack"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Registration started.")

	builder := registration.NewPayloadBuilder(registration.NewFileReader(), os.Getenv("SLACK_CONNECTOR_NAME"), os.Getenv("SLACK_BOT_TOKEN"))
	id, err := registration.NewApplicationRegistryClient(builder, 5, 10).RegisterService()
	if err != nil {
		log.Fatal("Fatal error: ", err.Error())
	}
	log.WithFields(log.Fields{
		"id": id,
	}).Info("Service registered.")

	kyma := events.NewSender(&http.Client{}, events.NewValidator(), "http://event-publish-service.kyma-system:8080/v1/events")
	webhook := handlers.NewWebHookHandler(
		slack.NewReceivingEventsWrapper(os.Getenv("SLACK_SECRET")),
		kyma,
	)

	http.HandleFunc("/webhook", webhook.HandleWebhook)
	log.Info(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
