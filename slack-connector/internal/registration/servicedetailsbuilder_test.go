package registration_test

import (
	"encoding/json"
	"testing"

	"github.com/kyma-incubator/hack-showcase/slack-connector/internal/apperrors"
	"github.com/kyma-incubator/hack-showcase/slack-connector/internal/registration"
	"github.com/kyma-incubator/hack-showcase/slack-connector/internal/registration/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBuildServiceDetails(t *testing.T) {
	t.Run("should return proper values", func(t *testing.T) {
		//given
		mockOSCommunicator := &mocks.OSCommunicator{}
		fileBody := []byte(`{"json":"value"}`)
		jsonBody := json.RawMessage(`{"json":"value"}`)
		mockOSCommunicator.On("ReadFile", "slackasyncapi.json").Return(fileBody, nil)
		mockOSCommunicator.On("GetEnv", "SLACK_CONNECTOR_NAME").Return("slack-connector")
		builder := registration.NewServiceDetailsBuilder(mockOSCommunicator)
		url := "https://raw.githubusercontent.com/kyma-incubator/hack-showcase/slack-connector-boilerplate/slack-connector/internal/registration/configs/slackopenapi.json"

		//when
		details, err := builder.BuildServiceDetails()

		//then
		assert.NoError(t, err)
		assert.Equal(t, "Kyma", details.Provider)
		assert.Equal(t, "Slack Connector, which is used for registering Slack API in Kyma", details.Description)
		assert.Equal(t, "https://slack.com/api/", details.API.TargetURL)
		assert.Equal(t, jsonBody, details.Events.Spec)
		assert.Equal(t, url, details.API.SpecificationURL)
	})

	t.Run("should return error and empty ServiceDetails{}", func(t *testing.T) {
		mockOSCommunicator := &mocks.OSCommunicator{}
		mockOSCommunicator.On("ReadFile", "slackasyncapi.json").Return(nil, apperrors.Internal("error"))
		mockOSCommunicator.On("GetEnv", "SLACK_CONNECTOR_NAME").Return("slack-connector")
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
		mockOSCommunicator.On("GetEnv", "SLACK_CONNECTOR_NAME").Return("slack-connector")
		targetURL := "http://application-registry-external-api.kyma-integration.svc.cluster.local:8081/slack-connector-app/v1/metadata/services"
		builder := registration.NewServiceDetailsBuilder(mockOSCommunicator)

		//when
		path := builder.GetApplicationRegistryURL()

		//then
		assert.Equal(t, targetURL, path)
	})
}
