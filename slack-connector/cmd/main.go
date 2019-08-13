package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("server starter")

	id, err := registration.RegisterService()
	if err != nil {
		log.Fatal("Fatal error: ", err.Error())
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info("Service registered")

}
