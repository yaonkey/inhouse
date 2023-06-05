package main

import (
	"log"
	"net/http"
	"yaonkey/inhouse/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)

	log.Printf("Server running on http://127.0.0.1:8989/")
	log.Fatal(http.ListenAndServe(":8989", r))
}
