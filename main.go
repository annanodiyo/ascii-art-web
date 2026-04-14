package main

import (
	"log"
	"net/http"

	"ascii-art-web/handlers"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/ascii", handlers.AsciiHandler)

	log.Println("Server starting on :8080...")

	// 👇 Wrap DefaultServeMux with middleware
	err := http.ListenAndServe(":8080", corsMiddleware(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}
}
