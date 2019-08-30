package hook_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/hook"
	"github.com/stretchr/testify/assert"
)

const sampleToken = "1234-567-890"

func exampleHookCreate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}

func exampleHookUnprocessableEntity(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnprocessableEntity)
}

func TestCreate(t *testing.T) {
	t.Run("should return nil", func(t *testing.T) {
		//given
		handler := http.HandlerFunc(exampleHookCreate)
		server := httptest.NewServer(handler)
		defer server.Close()
		creator := hook.NewHook("URL")
		//when
		err := creator.Create(sampleToken, server.URL)
		//then
		assert.NoError(t, err)
	})

	t.Run("should return error", func(t *testing.T) {
		//given
		handler := http.HandlerFunc(exampleHookUnprocessableEntity)
		server := httptest.NewServer(handler)
		defer server.Close()
		creator := hook.NewHook("URL")
		//when
		err := creator.Create(sampleToken, server.URL)
		//then
		assert.Error(t, err)
	})
}
