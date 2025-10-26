package main

import (
	"Copycat/internal/web"
	"net/http"
)

func main() {
	router := web.SetUpRoutes()

	http.ListenAndServe(":8080", router)
}
