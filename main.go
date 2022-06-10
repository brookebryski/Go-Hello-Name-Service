package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func run() {
	router := httprouter.New()
	router.ServeFiles("/js/*filepath", http.Dir("js"))
	router.ServeFiles("/css/*filepath", http.Dir("css"))

	router.GET("/", serveHomepage)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	run()
}
