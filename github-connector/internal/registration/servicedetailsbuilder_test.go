package registration_test

import (
	"encoding/json"
	"testing"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/apperrors"
	"github.com/kyma-incubator/hack-showcase/github-connector/internal/registration"
	"github.com/kyma-incubator/hack-showcase/github-connector/internal/registration/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBuildServiceDetails(t *testing.T) {
	t.Run("should return proper values", func(t *testing.T) {
		//given
		mockOSCommunicator := &mocks.OSCommunicator{}
		fileBody := []byte(`{"json":"value"}`)
		jsonBody := json.RawMessage(`{"json":"value"}`)
		mockOSCommunicator.On("ReadFile", "githubasyncapi.json").Return(fileBody, nil)
		mockOSCommunicator.On("GetEnv", "GITHUB_CONNECTOR_NAME").Return("github-connector")
		builder := registration.NewServiceDetailsBuilder(mockOSCommunicator)
		url := "https://raw.githubusercontent.com/colunira/github-openapi/master/githubopenAPI.json"

		//when
		details, err := builder.BuildServiceDetails()

		//then
		assert.NoError(t, err)
		assert.Equal(t, "Kyma", details.Provider)
		assert.Equal(t, "GitHub Connector, which can be used for communication and handling events from GitHub", details.Description)
		assert.Equal(t, "https://api.github.com", details.API.TargetURL)
		assert.Equal(t, jsonBody, details.Events.Spec)
		assert.Equal(t, url, details.API.SpecificationURL)
	})

	t.Run("should return error and empty ServiceDetails{}", func(t *testing.T) {
		mockOSCommunicator := &mocks.OSCommunicator{}
		mockOSCommunicator.On("ReadFile", "githubasyncapi.json").Return(nil, apperrors.Internal("error"))
		mockOSCommunicator.On("GetEnv", "GITHUB_CONNECTOR_NAME").Return("github-connector")
		builder := registration.NewServiceDetailsBuilder(mockOSCommunicator)

		//when
		details, err := builder.BuildServiceDetails()

		//then
		assert.Error(t, err)
		assert.Equal(t, registration.ServiceDetails{}, details)
	})
}

func TestGetApplicationRegistryURL(t *testing.T) {
	t.Run("should return proper URL", func(t *testing.T) {
		//given
		mockOSCommunicator := &mocks.OSCommunicator{}
		mockOSCommunicator.On("GetEnv", "GITHUB_CONNECTOR_NAME").Return("github-connector")
		targetURL := "http://application-registry-external-api.kyma-integration.svc.cluster.local:8081/github-connector-app/v1/metadata/services"
		builder := registration.NewServiceDetailsBuilder(mockOSCommunicator)

		//when
		path := builder.GetApplicationRegistryURL()

		//then
		assert.Equal(t, targetURL, path)
	})
}
