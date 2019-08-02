package kymasender

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/apperrors"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/eventparser"
	"github.com/kyma-incubator/hack-showcase/github-connector/internal/eventparser/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type toJSON struct {
	TestJSON string `json:TestJSON"`
}

type ClientMock struct {
}

func (c *ClientMock) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{}, nil
}

func TestSendToKyma_ProperArguments(t *testing.T) {
	t.Run("should return no error", func(t *testing.T) {
		k := NewKymaSenderWrapper(&ClientMock{}, eventparser.NewEventParser())
		payload := toJSON{TestJSON: "test"}
		toSend, err := json.Marshal(payload)
		require.NoError(t, err)
		assert.Equal(t, nil, k.SendToKyma("issuesevent.opened", "v1", "", "github-connector-app", json.RawMessage(toSend)))
	})
}

func TestWebhookHandler_sendToKyma(t *testing.T) {
	t.Run("should return no error", func(t *testing.T) {
		payload := toJSON{TestJSON: "test"}
		toSend, err := json.Marshal(payload)
		require.NoError(t, err)
		mockParser := &mocks.EventParser{}
		mockParser.On("GetEventRequestPayload", "", "v1", "", "github-connector-app", json.RawMessage(toSend)).Return(eventparser.EventRequestPayload{}, apperrors.Internal("test"))
		k := NewKymaSenderWrapper(&ClientMock{}, mockParser)
		expected := apperrors.Internal("test")
		actual := k.SendToKyma("", "v1", "", "github-connector-app", json.RawMessage(toSend))
		assert.Equal(t, expected.Code(), actual.Code())
	})
}
