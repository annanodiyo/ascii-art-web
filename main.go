package main

import (
	"net/http"

	"ascii-art-web/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	// http.HandleFunc("ascii-art",handlers.AsciiHandler)

	http.ListenAndServe(":8080", nil)
}
