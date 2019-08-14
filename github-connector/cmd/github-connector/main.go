package main

import (
	"net/http"
	"os"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/github"
	"github.com/kyma-incubator/hack-showcase/github-connector/internal/registration"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/events"
	"github.com/kyma-incubator/hack-showcase/github-connector/internal/handlers"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("server started")

	id, err := registration.RegisterService()
	if err != nil {
		log.Fatal("Fatal error: ", err.Error())
	}
	log.WithFields(log.Fields{
		"id": id,
	}).Info("Service registered")

	kyma := events.NewSender(&http.Client{}, events.NewValidator(), "http://event-bus-publish.kyma-system:8080/v1/events")
	webhook := handlers.NewWebHookHandler(
		github.ReceivingEventsWrapper{},
		kyma,
	)

	http.HandleFunc("/webhook", webhook.HandleWebhook)
	log.Info(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
