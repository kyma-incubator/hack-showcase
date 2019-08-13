package main

import (
	"github.com/kyma-incubator/hack-showcase/slack-connector/internal/registration"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("server started")

	builder := registration.NewServiceDetailsBuilder()
	requestSender := registration.NewRegisterRequestSender()
	service := registration.NewServiceRegister(builder, requestSender)
	id, err := service.RegisterService()
	if err != nil {
		log.Fatal("Fatal error: ", err.Error())
	}
	log.WithFields(log.Fields{
		"id": id,
	}).Info("Service registered")
}
