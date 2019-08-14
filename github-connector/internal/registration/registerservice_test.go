package registration_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/apperrors"

	"github.com/stretchr/testify/assert"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/registration"
	"github.com/kyma-incubator/hack-showcase/github-connector/internal/registration/mocks"
)

const expectedID = "123-456-789"

func exampleServiceID(w http.ResponseWriter, r *http.Request) {
	id := registration.RegisterResponse{ID: expectedID}
	res, err := json.Marshal(id)
	if err != nil {
	}
	w.Write(res)
}

func TestRegisterService(t *testing.T) {
	t.Run("should return service ID", func(t *testing.T) {
		//given
		handler := http.HandlerFunc(exampleServiceID)
		server := httptest.NewServer(handler)
		defer server.Close()

		mockBuilder := &mocks.Builder{}
		mockBuilder.On("BuildServiceDetails").Return(registration.ServiceDetails{}, nil)
		mockBuilder.On("GetApplicationRegistryURL").Return(server.URL)

		service := registration.NewServiceRegister("deploymentEnvName", mockBuilder)

		//when
		id, err := service.RegisterService()

		//then
		assert.NoError(t, err)
		assert.Equal(t, expectedID, id)

	})

	t.Run("should return an error when cannot build Service Details", func(t *testing.T) {
		//given
		handler := http.HandlerFunc(exampleServiceID)
		server := httptest.NewServer(handler)
		defer server.Close()

		mockBuilder := &mocks.Builder{}
		mockBuilder.On("BuildServiceDetails").Return(registration.ServiceDetails{}, apperrors.Internal("error"))
		mockBuilder.On("GetApplicationRegistryURL").Return(server.URL)

		service := registration.NewServiceRegister("deploymentEnvName", mockBuilder)

		//when
		id, err := service.RegisterService()

		//then
		assert.Error(t, err)
		assert.EqualError(t, err, "While building service details json: error")
		assert.Equal(t, "", id)
	})

	// t.Run("should return an error when cannot reach server", func(t *testing.T) {
	// 	//given
	// 	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		w.WriteHeader(500)
	// 	})
	// 	server := httptest.NewServer(handler)
	// 	defer server.Close()

	// 	mockBuilder := &mocks.Builder{}
	// 	mockBuilder.On("BuildServiceDetails").Return(registration.ServiceDetails{}, nil)
	// 	mockBuilder.On("GetApplicationRegistryURL").Return(server.URL)

	// 	service := registration.NewServiceRegister("deploymentEnvName", mockBuilder)

	// 	//when
	// 	id, err := service.RegisterService()

	// 	//then
	// 	assert.Error(t, err)
	// 	assert.EqualError(t, err, "While trying to register service: error")
	// 	assert.Equal(t, "", id)
	// })
}
