package registration_test

import (
	"encoding/json"
	"testing"

	"github.com/kyma-incubator/github-slack-connectors/slack-connector/internal/apperrors"
	"github.com/kyma-incubator/github-slack-connectors/slack-connector/internal/registration"
	"github.com/kyma-incubator/github-slack-connectors/slack-connector/internal/registration/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBuild(t *testing.T) {
	t.Run("should return proper values", func(t *testing.T) {
		//given
		mockFileReader := &mocks.FileReader{}
		fileBody := []byte(`{"json":"value"}`)
		jsonBody := json.RawMessage(`{"json":"value"}`)
		mockFileReader.On("Read", "slackasyncapi.json").Return(fileBody, nil)
		builder := registration.NewPayloadBuilder(mockFileReader, "slack-connector", "token")
		url := registration.SpecificationURL

		//when
		details, err := builder.Build()

		//then
		assert.NoError(t, err)
		assert.Equal(t, "Kyma", details.Provider)
		assert.Equal(t, "Slack Connector, which is used for registering Slack API in Kyma", details.Description)
		assert.Equal(t, "https://slack.com/api", details.API.TargetURL)
		assert.Equal(t, jsonBody, details.Events.Spec)
		assert.Equal(t, url, details.API.SpecificationURL)
	})

	t.Run("should return error and empty ServiceDetails{}", func(t *testing.T) {
		mockFileReader := &mocks.FileReader{}
		mockFileReader.On("Read", "slackasyncapi.json").Return(nil, apperrors.Internal("error"))
		builder := registration.NewPayloadBuilder(mockFileReader, "slack-connector", "token")

		//when
		details, err := builder.Build()

		//then
		assert.Error(t, err)
		assert.Equal(t, registration.ServiceDetails{}, details)
	})
}

func TestGetApplicationRegistryURL(t *testing.T) {
	t.Run("should return proper URL", func(t *testing.T) {
		//given
		mockFileReader := &mocks.FileReader{}
		targetURL := "http://application-registry-external-api.kyma-integration.svc.cluster.local:8081/slack-connector-app/v1/metadata/services"
		builder := registration.NewPayloadBuilder(mockFileReader, "slack-connector", "token")

		//when
		path := builder.GetApplicationRegistryURL()

		//then
		assert.Equal(t, targetURL, path)
	})
}
