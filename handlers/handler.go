package handlers

import (
	"html/template"
	"net/http"
	"strings"

	"ascii-art-web/ascii"
)

// Load template once
var tmpl = template.Must(template.ParseFiles("templates/index.html"))

// Struct sent to HTML
type Response struct {
	Result string
	Error  string
	Text   string
}

// GET "/"
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	tmpl.Execute(w, Response{
		Text: "TRy it out",
	})
}

// POST "/ascii"
func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	// Only POST allowed
	if r.Method != http.MethodPost {
		http.Error(w, "400 bad request", http.StatusBadRequest)
		return
	}

	text := strings.TrimSpace(r.FormValue("text"))
	banner := r.FormValue("banner")

	// Validation
	if text == "" {
		w.WriteHeader(http.StatusBadRequest)
		tmpl.Execute(w, Response{Error: "Text cannot be empty"})
		return
	}

	if banner == "" {
		w.WriteHeader(http.StatusBadRequest)
		tmpl.Execute(w, Response{Error: "Please select a banner"})
		return
	}

	var result string
	var err error

	// NEW FEATURE: handle "all"
	if banner == "all" {
		result, err = ascii.AllSamples(text)
	} else {
		result, err = ascii.Generate(text, banner)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tmpl.Execute(w, Response{Error: err.Error()})
		return
	}

	tmpl.Execute(w, Response{
		Result: result,
		Text:   text,
	})
}
