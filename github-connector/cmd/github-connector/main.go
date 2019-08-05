package main

import (
	"net/http"
	"os"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/eventparser"

	registerservice "github.com/kyma-incubator/hack-showcase/github-connector/internal/application_registry"
	"github.com/kyma-incubator/hack-showcase/github-connector/internal/githubwrappers"
	"github.com/kyma-incubator/hack-showcase/github-connector/internal/handlers"
	"github.com/kyma-incubator/hack-showcase/github-connector/internal/kymasender"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("server started")

	id, err := registerservice.RegisterService()
	if err != nil {
		log.Fatal("Fatal error: ", err.Error())
	}
	log.WithFields(log.Fields{
		"id": id,
	}).Info("Service registered")

	kyma := kymasender.NewWrapper(&http.Client{}, eventparser.NewEventParser())
	webhook := handlers.NewWebHookHandler(
		githubwrappers.ReceivingEventsWrapper{},
		kyma,
	)

	http.HandleFunc("/webhook", webhook.HandleWebhook)
	log.Info(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
