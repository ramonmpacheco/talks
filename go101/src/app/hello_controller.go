package app

import (
	"net/http"

	"github.com/go-chi/render"
)

type helloController struct{}

func NewHelloController() *helloController {
	return &helloController{}
}

func (g *helloController) greetings(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, "Hello Talks!!!")
}
