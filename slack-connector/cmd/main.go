package main

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("server started")

	id, err := registration.RegisterService()
	if err != nil {
		log.Fatal("Fatal error: ", err.Error())
	}
	log.WithFields(log.Fields{"id": id,}).Info("Service registered")
}
