package main

import (
	"log"
	"net/http"
	"sever-client/src/router"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	router.Register(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
