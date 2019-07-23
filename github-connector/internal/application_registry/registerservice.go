package registerservice

import (
	"log"
	"time"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/model"
)

var jsonBody = model.ServiceDetails{
	Provider:    "kyma",
	Name:        "github-connector",
	Description: "Boilerplate for GitHub connector",
	API: &model.API{
		TargetURL:        "https://api.github.com",
		SpecificationURL: "https://raw.githubusercontent.com/colunira/github-openapi/master/githubopenAPI.yaml",
	},
}

var url = "http://application-registry-external-api.kyma-integration.svc.cluster.local:8081/github-connector/v1/metadata/services"

//RegisterService - register service in Kyma and get a response
func RegisterService() {

	var id string
	var err error
	for i := 0; i < 10; i++ {
		id, err = SendRegisterRequest(jsonBody, url)
		if err == nil {
			break
		}

		time.Sleep(5 * time.Second)
	}

	if err != nil {
		panic(err)
	}

	log.Printf("Application ID: " + id)
}
