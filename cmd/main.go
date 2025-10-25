package main

import (
	"Copycat/pkg/app"
	"net/http"
)

func main() {
	router := app.SetUpRoutes()

	http.ListenAndServe(":8080", router)
}
