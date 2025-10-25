package routes

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UserRoutes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		newUrl := fmt.Sprintf("/user/%d", 1)
		http.Redirect(w, r, newUrl, http.StatusPermanentRedirect)
	})

	router.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		userId := chi.URLParam(r, "id")

		if userId != "1" {
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
		}

		w.Write([]byte("User page"))
	})

	router.Get("/edit", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Edit page"))
	})

	return router
}
