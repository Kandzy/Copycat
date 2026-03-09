package main

import (
	"Copycat/internal/web"
	"Copycat/internal/web/environment"
	"flag"
	"net/http"
)

func main() {
	environment.DevMode = flag.Bool("dev", false, "Run in development mode")
	flag.Parse()

	router := web.SetUpRoutes()

	http.ListenAndServe(":8080", router)
}
