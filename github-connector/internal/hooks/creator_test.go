package hooks_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/hooks"
	"github.com/stretchr/testify/assert"
)

func exampleHookCreate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}

func TestCreate(t *testing.T) {
	t.Run("should return nil", func(t *testing.T) {
		//given
		handler := http.HandlerFunc(exampleHookCreate)
		server := httptest.NewServer(handler)
		defer server.Close()
		creator := hooks.NewCreator("token", server.URL)
		//when
		err := creator.Create("URL")
		//then
		assert.NoError(t, err)
	})
}
