package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<h1>Hello from Go backend!</h1><p>This is rendered by Go.</p>")
	})

	http.ListenAndServe(":8080", nil)
}
