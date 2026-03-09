package web

import (
	"Copycat/internal/web/assets"
	"Copycat/internal/web/routes"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetUpRoutes() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	staticFS := assets.GetStaticFS()
	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(staticFS)))

	router.Mount("/", routes.HomepageRoutes())
	router.Mount("/login", routes.LoginPage())
	router.Mount("/signup", routes.SignupPage())
	router.Mount("/admin", routes.AdminPageRoutes())
	router.Mount("/user", routes.UserRoutes())

	return router
}
