package main

import (
	"log"
	"net/http"
	"os"

	"github.com/google/go-github/github"
)

//WebHookStruct that bundles the github library functions into one struct with a Validator interface
type WebHookStruct struct {
}

//ValidatePayload is a function used for checking whether the secret provided in the request is correct
func (wh WebHookStruct) ValidatePayload(r *http.Request, b []byte) ([]byte, error) {
	return github.ValidatePayload(r, b)
}

//ParseWebHook parses the raw json payload into an event struct
func (wh WebHookStruct) ParseWebHook(s string, b []byte) (interface{}, error) {
	return github.ParseWebHook(s, b)
}

//GetToken is a function that looks for the secret in the environment
func (wh WebHookStruct) GetToken() string {
	return os.Getenv("GITHUB-TOKEN")
}

func main() {

	log.Println("server started")
	str := WebHookStruct{}
	wh := NewWebHookHandler(str)

	http.HandleFunc("/webhook", wh.handleWebhook)
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), nil))
}
