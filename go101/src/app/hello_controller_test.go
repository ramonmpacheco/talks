package app_test

import (
	"go-101/app"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello_greetings(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello", nil)
	assert.Nil(t, err)

	resp := httptest.NewRecorder()

	app.InitRoutes().ServeHTTP(resp, req)

	assert.EqualValues(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Hello Talks!!!")
}
