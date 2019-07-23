package main

import (
	"net/http"
	"os"

	tmpfunc "github.com/kyma-incubator/hack-showcase/github-connector/internal"
	"github.com/kyma-incubator/hack-showcase/github-connector/internal/githubwrappers"
	"github.com/kyma-incubator/hack-showcase/github-connector/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {

	tmpfunc.TmpFunc()

	log.Println("server started")
	webhook := handlers.NewWebHookHandler(
		githubwrappers.ReceivingEventsWrapper{},
	)

	http.HandleFunc("/webhook", webhook.HandleWebhook)
	log.Info(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
