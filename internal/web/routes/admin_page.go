package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func AdminPageRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Admin dashboard"))
	})

	r.Get("/stats", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Admin stats page"))
	})

	return r
}
