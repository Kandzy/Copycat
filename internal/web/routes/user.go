package routes

import (
	"Copycat/internal/web/assets"
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

		tpl := assets.GetTemplates("user")

		err := tpl.ExecuteTemplate(w, "details.html", map[string]string{
			"UserId":    userId,
			"Username":  "Kandzy",
			"DiscordId": "Kandzy#1234",
			"UserGroup": "Admin, Moderator",
		})
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			fmt.Println("Error rendering template:", err)
			return
		}
	})

	router.Get("/edit", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Edit page"))
	})

	return router
}
