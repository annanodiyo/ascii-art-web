package handlers

import (
	"html/template"
	"net/http"

	"ascii-art-web/ascii"
	// "ascii-art-web/ascii"
)

var tmplate = template.Must(template.ParseFiles("templates/index.html"))

type Response struct {
	Result string
	Error  string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	tmplate.Execute(w, nil)
}

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	// allow only post methods
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
	result, err := ascii.Generate(text, banner)
	if err != nil {
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}
	tmplate.Execute(w, Response{Result: result})
}
