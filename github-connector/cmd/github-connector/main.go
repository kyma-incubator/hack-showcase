package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/github"
	"github.com/kyma-incubator/hack-showcase/github-connector/internal/hook"
	"github.com/kyma-incubator/hack-showcase/github-connector/internal/registration"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/events"
	"github.com/kyma-incubator/hack-showcase/github-connector/internal/handlers"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Registration started.")
	receive := flag.Bool("receiving", true, "Specifies if Connector should subscribe events from GitHub")
	send := flag.Bool("sending", true, "Specifies if Connector should send events to GitHub")
	flag.Parse()
	log.Infof("Events receiving: %t", *receive)
	log.Infof("Events sending: %t", *send)

	builder := registration.NewPayloadBuilder(registration.NewFileReader(), os.Getenv("GITHUB_CONNECTOR_NAME"), os.Getenv("GITHUB_TOKEN"), *receive, *send)
	id, err := registration.NewApplicationRegistryClient(builder, 5, 10).RegisterService()

	if err != nil {
		log.Fatal("Fatal error: ", err.Error())
	}
	log.WithFields(log.Fields{
		"id": id,
	}).Info("Service registered")

	if *receive {
		webHook := hook.NewHook(os.Getenv("KYMA_ADDRESS"))
		hookSecret, err := webHook.Create(os.Getenv("GITHUB_TOKEN"), os.Getenv("GITHUB_REPO_URL"))
		if err != nil {
			log.Fatalf("Fatal error: %s", err.Error())
		}
		log.Info("Webhook created!")

		kyma := events.NewSender(&http.Client{}, events.NewValidator(), "http://event-publish-service.kyma-system:8080/v1/events")
		webhook := handlers.NewWebHookHandler(
			github.NewReceivingEventsWrapper(hookSecret),
			kyma,
		)

		http.HandleFunc("/webhook", webhook.HandleWebhook)
		log.Info(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
	} else {
		log.Info("Waiting for input...")
		for {
		}
	}
}
