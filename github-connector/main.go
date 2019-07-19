package main

import (
	"log"
	"net/http"
	"os"

	"github.com/google/go-github/github"
)

type WebhookStructHelper struct {
}

func (wh WebhookStructHelper) ValidatePayload(r *http.Request, b []byte) ([]byte, error) {
	return github.ValidatePayload(r, b)
}

func (wh WebhookStructHelper) ParseWebHook(s string, b []byte) (interface{}, error) {
	return github.ParseWebHook(s, b)
}

func (wh WebhookStructHelper) GetToken() string {
	return "test"
}

func main() {

	log.Println("server started")
	str := WebhookStructHelper{}
	wh := NewWebhookHandler(str)

	http.HandleFunc("/webhook", wh.handleWebhook)
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), nil))
}
