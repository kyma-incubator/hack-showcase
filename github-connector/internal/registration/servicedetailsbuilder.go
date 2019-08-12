package registration

import (
	"io/ioutil"
	"os"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/apperrors"
)

//Builder is an interface containing all necessary functions required to build an ServiceDetails structure
type Builder interface {
	BuildServiceDetails(string, string) (ServiceDetails, error)
}

//ServiceDetailsBuilder is used for mocking building ServiceDetails struct
type ServiceDetailsBuilder struct {
	builder Builder
}

var jsonBody = ServiceDetails{
	Provider:    "Kyma",
	Name:        os.Getenv("GITHUB_CONNECTOR_NAME"),
	Description: "GitHub Connector, which can be used for communication and handling events from GitHub",
	API: &API{
		TargetURL: "https://api.github.com",
	},
}

//BuildServiceDetails creates a ServiceDetails structure with provided API specification URL and events description json path
func BuildServiceDetails(url string, path string) (ServiceDetails, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return ServiceDetails{}, apperrors.Internal("While reading file: %s", err)
	}
	data := []byte(file)

	jsonBody.Events.Spec = data
	jsonBody.API.SpecificationURL = url
	return jsonBody, nil
}
