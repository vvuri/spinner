package main

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

const (
	baseUrl = "http://localhost:8099"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Get("")

	http.ListenAndServe(baseUrl, r)
}
