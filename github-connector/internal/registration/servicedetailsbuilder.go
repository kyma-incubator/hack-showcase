package registration

import (
	"io/ioutil"
	"os"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/apperrors"
)

const specificationURL = "https://raw.githubusercontent.com/colunira/github-openapi/master/githubopenAPI.json"

//Builder is an interface containing all necessary functions required to build an ServiceDetails structure
type Builder interface {
	BuildServiceDetails() (ServiceDetails, error)
}

//ServiceDetailsBuilder is used for mocking building ServiceDetails struct
type serviceDetailsBuilder struct {
	builder Builder
}

//NewServiceDetailsBuilder creates a serviceDetailsBuilder instance
func NewServiceDetailsBuilder() serviceDetailsBuilder {
	return serviceDetailsBuilder{}
}

//BuildServiceDetails creates a ServiceDetails structure with provided API specification URL
func (r serviceDetailsBuilder) BuildServiceDetails() (ServiceDetails, error) {

	var jsonBody = ServiceDetails{
		Provider:    "Kyma",
		Name:        os.Getenv("GITHUB_CONNECTOR_NAME"),
		Description: "GitHub Connector, which can be used for communication and handling events from GitHub",
		API: &API{
			TargetURL: "https://api.github.com",
		},
	}
	file, err := ioutil.ReadFile("githubasyncapi.json")
	if err != nil {
		return ServiceDetails{}, apperrors.Internal("While reading githubasyncapi.json: %s", err)
	}
	jsonBody.Events = &Events{Spec: file}

	jsonBody.API.SpecificationURL = specificationURL
	return jsonBody, nil
}
