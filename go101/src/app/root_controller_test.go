package app_test

import (
	"encoding/json"
	"go-101/app"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoot_showRoutes(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)

	resp := httptest.NewRecorder()

	app.InitRoutes().ServeHTTP(resp, req)

	var body interface{}
	json.NewDecoder(resp.Body).Decode(&body)

	assert.EqualValues(t, "Root 😎", body.(map[string]interface{})["name"])
}
