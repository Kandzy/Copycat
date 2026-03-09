package api

import (
	"Copycat/internal/api/routes/discord"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetUpRoutes() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Mount("/discord/command", discord.CommandRoutes())

	return router
}
