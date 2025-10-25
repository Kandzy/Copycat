package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func LoginPage() chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Login Page"))
	})

	return r
}
