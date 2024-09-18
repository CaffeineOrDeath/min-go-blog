package main

import (
	"net/http"

	"mini-go-blog/internal/views"
	"github.com/a-h/templ"
)

func main() {
	port := ":8080"
	component := views.Layout("Home")
    
	http.Handle("/", templ.Handler(component))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	
	http.ListenAndServe(port, nil)
}
