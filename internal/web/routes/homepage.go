package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HomepageRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Copycat homepage"))
	})

	r.Get("/about", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Copycat about page"))
	})

	return r
}
