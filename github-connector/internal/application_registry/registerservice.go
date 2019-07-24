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
		SpecificationURL: "https://raw.githubusercontent.com/colunira/github-openapi/master/githubopenAPI.json",
	},
	Events: &model.Events{
		Spec: []byte(`{
			"asyncapi": "1.0.0",
			"info": {
				"title": "github-events",
				"version": "v1",
				"description": "Github Events v1"
			},
			"topics": {
				"issuesevent.opened.v1": {
					"subscribe": {
						"summary": "Github issue commented event v1",
						"payload": {
							"type": "object",
							"required": [
								"action"
							],
							"properties": {
								"action": {
									"type": "string",
									"example": "edited",
									"description": "The action that was performed.",
									"title": "Action"
								},
								"issue": {
									"type": "object"
								},
								"changes": {
									"type": "object"
								},
								"repository": {
									"type": "object"
								},
								"sender": {
									"type": "object"
								}
							}
						}
					}
				}
			}
		}`),
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
