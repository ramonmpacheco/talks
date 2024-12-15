package app

import (
	"net/http"

	"github.com/go-chi/render"
)

type rootController struct{}

func NewRootController() *rootController {
	return &rootController{}
}

func (g *rootController) root(w http.ResponseWriter, r *http.Request) {
	response := NewRootResponse("Root 😎")
	response.appenPath("hello", "Greetings page", []string{"/hello"})
	response.appenPath(
		"counter", "Use to be aware of the word size", []string{"/counter?word=yourWord"},
	)

	render.Status(r, http.StatusOK)
	render.JSON(w, r, response)
}
