package handlers

import (
	"ascii-art-web/ascii"
	"html/template"
	"net/http"
)

// GET /
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

// POST /ascii-art
func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "400 bad request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" || banner == "" {
		http.Error(w, "400 bad request", http.StatusBadRequest)
		return
	}

	result, err := ascii.GenerateAscii(text, banner)
	if err != nil {
		if err.Error() == "banner not found" {
			http.Error(w, "404 banner not found", http.StatusNotFound)
			return
		}
		if err.Error() == "invalid character" {
			http.Error(w, "400 bad request", http.StatusBadRequest)
			return
		}
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/result.html")
	if err != nil {
		http.Error(w, "404 template not found", http.StatusNotFound)
		return
	}

	tmpl.Execute(w, result)
}