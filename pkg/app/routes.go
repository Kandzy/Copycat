package app

import (
	"Copycat/pkg/app/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetUpRoutes() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Mount("/", routes.HomepageRoutes())
	router.Mount("/login", routes.LoginPage())
	router.Mount("/admin", routes.AdminPageRoutes())
	router.Mount("/user", routes.UserRoutes())

	return router
}
