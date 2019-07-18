package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	log.Println("server started")

	wh := NewWebhookHandler(nil)

	http.HandleFunc("/webhook", wh.handleWebhook)
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), nil))
}
