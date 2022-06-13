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
	router.POST("/get_name", getName)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	run()
}
