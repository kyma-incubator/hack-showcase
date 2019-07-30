package httperrors

import (
	"encoding/json"
	"net/http"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/apperrors"
	log "github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

func AppErrorToResponse(appError apperrors.AppError, detailedErrorResponse bool) (status int, body ErrorResponse) {
	httpCode := errorCodeToHttpStatus(appError.Code())
	errorMessage := appError.Error()
	return formatErrorResponse(httpCode, errorMessage, detailedErrorResponse)
}

func errorCodeToHttpStatus(code int) int {
	switch code {
	case apperrors.CodeInternal:
		return http.StatusInternalServerError
	case apperrors.CodeNotFound:
		return http.StatusNotFound
	case apperrors.CodeAlreadyExists:
		return http.StatusConflict
	case apperrors.CodeWrongInput:
		return http.StatusBadRequest
	case apperrors.CodeUpstreamServerCallFailed:
		return http.StatusBadGateway
	case apperrors.CodeAuthenticationFailed:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}

func formatErrorResponse(httpCode int, errorMessage string, detailedErrorResponse bool) (status int, body ErrorResponse) {
	if isInternalError(httpCode) && !detailedErrorResponse {
		return httpCode, ErrorResponse{httpCode, "Internal error."}
	}
	return httpCode, ErrorResponse{httpCode, errorMessage}
}

func isInternalError(httpCode int) bool {
	return httpCode == http.StatusInternalServerError
}

//SendErrorResponse prepares the http error response and sends it to the client
func SendErrorResponse(apperrptr *apperrors.AppError, wptr *http.ResponseWriter) {
	apperr := *apperrptr
	w := *wptr

	httpcode, resp := AppErrorToResponse(apperr, false)

	w.WriteHeader(httpcode)
	respJSON, err := json.Marshal(resp)

	if err != nil {
		log.Warn(apperrors.Internal("Failed to marshal error response: %s \n\tError body: %s", err, apperr.Error()))
		return
	}
	w.Write(respJSON)
	return
}
