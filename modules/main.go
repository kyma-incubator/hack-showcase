package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kyma-incubator/hack-showcase/modules/internal/eventparser"
)


//	sampleGithubPayload = [123 34 101 118 101 110 116 45 116 121 112 101 34 58 34 115 97 109 112 108 101 45 101 118 101 110 116 45 116 121 112 101 34 44 34 101 118 101 110 116 45 116 121 112 101 45 118 101 114 115 105 111 110 34 58 34 118 49 34 44 34 101 118 101 110 116 45 105 100 34 58 34 34 44 34 101 118 101 110 116 45 116 105 109 101 34 58 34 50 48 49 57 45 48 55 45 49 56 84 49 51 58 49 53 58 51 55 43 48 50 58 48 48 34 44 34 100 97 116 97 34 58 123 125 125] byte


func main() {
	fmt.Println(eventparser.GetEventRequestPayload("sample-event-type", "v1", "", make([]byte, 5)))
	r, err := json.Marshal(eventparser.GetEventRequestPayload("sample-event-type", "v1", "", make([]byte, 5)))
	if err != nil {
		fmt.Println("error: ", err)
	}
	os.Stdout.Write(r)
}
