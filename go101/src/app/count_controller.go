package app

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

type counterController struct{}

func NewCounterController() *counterController {
	return &counterController{}
}

func (g *counterController) count(w http.ResponseWriter, r *http.Request) {
	word := r.URL.Query().Get("word")

	if word == "" {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, "Word to count not found 🥲")
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, fmt.Sprintf("The word: %s has the size of %d", word, len(word)))
}
