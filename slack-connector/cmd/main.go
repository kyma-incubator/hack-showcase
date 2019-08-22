package main

import (
	"os"

	"github.com/kyma-incubator/hack-showcase/slack-connector/internal/registration"
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

	for true { // Temporary solution
	}
}
