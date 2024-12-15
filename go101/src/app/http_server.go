package app

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func InitRoutes() *chi.Mux {
	helloController := NewHelloController()
	counterController := NewCounterController()
	rootController := NewRootController()

	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Get("/", rootController.root)
		r.Get("/hello", helloController.greetings)
		r.Get("/counter", counterController.count)
	})

	return r
}

func Start(r *chi.Mux) {
	fmt.Println("About to start http server, port: 3000")
	http.ListenAndServe(":3000", r)
}
