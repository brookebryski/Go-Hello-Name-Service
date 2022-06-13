package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

type HomePage struct {
	Greeting string
}

type NameInput struct {
	Name string
}

func serveHomepage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var homepage HomePage
	homepage.Greeting = "Welcome"

	tmpl := template.Must(template.ParseFiles("html/homepage.html"))
	_ = tmpl.Execute(w, homepage)
}

func getName(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var data NameInput
	json.NewDecoder(r.Body).Decode(&data)
	fmt.Println(data.Name)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
	return
}
