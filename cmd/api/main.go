package main

import (
	"Copycat/internal/api"
	"net/http"
)

func main() {
	router := api.SetUpRoutes()

	http.ListenAndServe(":8080", router)
}
