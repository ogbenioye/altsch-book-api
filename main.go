package main

import (
	"log"
	"net/http"
	"sever-client/src/router"
	setup "sever-client/src/setups"

	"github.com/gorilla/mux"
)

func main() {
	setup.Connection()

	r := mux.NewRouter()
	router.RegisterRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
