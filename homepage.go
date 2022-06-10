package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

type HomePage struct {
	Name string
}

func serveHomepage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var homepage HomePage
	homepage.Name = "Brooke"

	tmpl := template.Must(template.ParseFiles("html/homepage.html"))
	_ = tmpl.Execute(w, homepage)
}
