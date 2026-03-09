package routes

import (
	"Copycat/internal/web/assets"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SignupPage() chi.Router {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tpl := assets.GetTemplates("")

		err := tpl.ExecuteTemplate(w, "signup.html", nil)

		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
	})

	return router
}
