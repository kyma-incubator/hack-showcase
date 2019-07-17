package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	log.Println("server started")
	http.HandleFunc("/webhook", handleWebhook)
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), nil))
}
