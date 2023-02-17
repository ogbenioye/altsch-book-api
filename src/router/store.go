package router

import (
	"sever-client/src/controller"

	"github.com/gorilla/mux"
)

var Register = func(router *mux.Router) {
	router.HandleFunc("/books", controller.CreateBook).Methods("POST")
}
