package tmpfunc

import (
	"log"
	"time"

	registerapp "github.com/kyma-incubator/hack-showcase/github-connector/internal/application_registry"
	"github.com/kyma-incubator/hack-showcase/github-connector/internal/model"
)

//TmpFunc - s
func TmpFunc() {

	// create json structure
	JSONBody := model.ServiceDetails{
		Provider:    "kyma",
		Name:        "github-connector",
		Description: "Boilerplate for GitHub connector",
		API: &model.API{
			TargetURL:        "https://console.35.195.62.81.xip.io/github-api",
			SpecificationURL: "https://raw.githubusercontent.com/colunira/github-openapi/master/githubopenAPI.yaml",
		},
	}
	url := "http://application-registry-external-api.kyma-integration.svc.cluster.local:8081/github-connector/v1/metadata/services"

	log.Println("===================")
	log.Println("====Register App===")
	log.Println("wait 30 sec")

	log.Println("try request")
	var id string
	var err error
	for i := 0; i < 10; i++ {
		log.Println("Try: " + string(i) + " time")

		id, err = registerapp.RegisterApp(JSONBody, url)
		if err == nil {
			break
		}

		time.Sleep(5 * time.Second)
	}

	if err != nil {
		panic(err)
	} else {
		log.Println(id)
		log.Println("Done! Wait 12h and die :)")
		log.Println("===================")
		time.Sleep(12 * time.Hour)
	}
}
