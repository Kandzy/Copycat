package discord

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func CommandRoutes() chi.Router {
	router := chi.NewRouter()

	router.Get("/userList", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("List of users"))
	})

	router.Get("/groupList", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("List of user groups"))
	})

	return router
}
