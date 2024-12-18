package app_test

import (
	"go-101/app"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounter_count(t *testing.T) {
	req, err := http.NewRequest("GET", "/count?word=teste", nil)
	assert.Nil(t, err)

	resp := httptest.NewRecorder()

	app.InitRoutes().ServeHTTP(resp, req)
	assert.EqualValues(t, http.StatusOK, resp.Code)

	assert.Contains(t, resp.Body.String(), "The word: teste has the size of 5")
}
