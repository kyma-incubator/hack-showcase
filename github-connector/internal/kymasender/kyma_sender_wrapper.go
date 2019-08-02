package kymasender

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/apperrors"
	"github.com/kyma-incubator/hack-showcase/github-connector/internal/eventparser"

	log "github.com/sirupsen/logrus"
)

type KymaSenderWrapper struct {
	parser eventparser.EventParser
	client HTTPClient
}

func NewKymaSenderWrapper(c HTTPClient, ep eventparser.EventParser) KymaSenderWrapper {
	return KymaSenderWrapper{client: c, parser: ep}
}

//HTTPClient is an interface use to allow mocking the http.Client methods
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func (k KymaSenderWrapper) SendToKyma(eventType, eventTypeVersion, eventID, sourceID string, data json.RawMessage) apperrors.AppError {
	toSend, apperr := k.parser.GetEventRequestPayload(eventType, eventTypeVersion, eventID, sourceID, data)
	if apperr != nil {
		return apperrors.Internal("While parsing the event payload: %s", apperr.Error())
	}
	jsonToSend, apperr := k.parser.GetEventRequestAsJSON(toSend)
	if apperr != nil {
		return apperrors.Internal("While getting the request as json %s", apperr.Error())
	}
	kymaRequest, err := http.NewRequest(http.MethodPost, "http://event-bus-publish.kyma-system:8080/v1/events",
		bytes.NewReader(jsonToSend))
	if err != nil {
		return apperrors.Internal("While creating an http request: %s", err.Error())
	}
	response, err := k.client.Do(kymaRequest)
	if err != nil {
		return apperrors.UpstreamServerCallFailed("While sending the event to the EventBus: %s", err.Error())
	}
	log.Info(response)
	return nil
}
