package main

import (
	"log"
	"net/http"
	"sever-client/src/router"
	setup "sever-client/src/setups"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	setup.Connection()

	r := mux.NewRouter()
	router.RegisterRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
