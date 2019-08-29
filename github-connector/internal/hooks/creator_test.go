package hooks_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/hooks"
	"github.com/stretchr/testify/assert"
)

const sampleToken = "1234-567-890"

func exampleHookCreate(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
	} else if string(body) == sampleToken {
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusUnprocessableEntity)
	}

}

func TestCreate(t *testing.T) {
	t.Run("should return nil", func(t *testing.T) {
		//given
		handler := http.HandlerFunc(exampleHookCreate)
		server := httptest.NewServer(handler)
		defer server.Close()
		creator := hooks.NewCreator(sampleToken, server.URL)
		//when
		err := creator.Create("URL")
		//then
		assert.NoError(t, err)
	})
}
