package events

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/apperrors"

	log "github.com/sirupsen/logrus"
)

//Wrapper is a struct used to allow mocking the SendToKyma function
type Wrapper struct {
	parser EventParser
	client HTTPClient
}

//NewWrapper is a function that creates new Wrapper with the passed in interfaces
func NewWrapper(c HTTPClient, ep EventParser) Wrapper {
	return Wrapper{client: c, parser: ep}
}

//HTTPClient is an interface use to allow mocking the http.Client methods
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

//SendToKyma is a function that sends the event given by the Github API to kyma's event bus
func (k Wrapper) SendToKyma(eventType, eventTypeVersion, eventID, sourceID string, data json.RawMessage) apperrors.AppError {
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
