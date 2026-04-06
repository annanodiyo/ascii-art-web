// package handlers

// import (
// 	"html/template"
// 	"net/http"
// )

// var tmplate = template.Must(template.ParseFiles("templates/index.html"))

// type Response struct {
// 	Result string
// 	Error  string
// }

// func HomeHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		http.Error(w, "404 not found", http.StatusNotFound)
// 		return
// 	}
// 	tmplate.Execute(w, nil)
// }
