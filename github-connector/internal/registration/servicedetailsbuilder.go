package registration

import (
	"os"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/registration/configs"
)

//Builder is an interface containing all necessary functions required to build an ServiceDetails structure
type Builder interface {
	BuildServiceDetails(string) (ServiceDetails, error)
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
func (r serviceDetailsBuilder) BuildServiceDetails(url string) (ServiceDetails, error) {

	var jsonBody = ServiceDetails{
		Provider:    "Kyma",
		Name:        os.Getenv("GITHUB_CONNECTOR_NAME"),
		Description: "GitHub Connector, which can be used for communication and handling events from GitHub",
		API: &API{
			TargetURL: "https://api.github.com",
		},
		Events: &Events{Spec: configs.GithubAsyncAPI},
	}

	jsonBody.API.SpecificationURL = url
	return jsonBody, nil
}
