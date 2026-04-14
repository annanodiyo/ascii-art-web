package handlers

import (
	"html/template"
	"net/http"

	"ascii-art-web/ascii"
)

// Load template once
var tmplate = template.Must(template.ParseFiles("templates/index.html"))

// Struct sent to HTML
type Response struct {
	Result string
	Error  string
}

// GET "/"
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	tmplate.Execute(w, nil)
}

// POST "/ascii"
func AsciiHandler(w http.ResponseWriter, r *http.Request) {

	// Only POST allowed
	if r.Method != http.MethodPost {
		http.Error(w, "400 bad request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	// Validation
	if text == "" {
		w.WriteHeader(http.StatusBadRequest)
		tmplate.Execute(w, Response{Error: "Text cannot be empty"})
		return
	}

	if banner == "" {
		w.WriteHeader(http.StatusBadRequest)
		tmplate.Execute(w, Response{Error: "Please select a banner"})
		return
	}

	var result string
	var err error

	// ✅ NEW FEATURE: handle "all"
	if banner == "all" {
		result, err = ascii.AllSamples(text)
	} else {
		result, err = ascii.Generate(text, banner)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tmplate.Execute(w, Response{Error: err.Error()})
		return
	}

	tmplate.Execute(w, Response{Result: result})
}