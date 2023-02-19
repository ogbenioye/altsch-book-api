package router

import (
	"sever-client/src/controller"

	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controller.CreateBook).Methods("POST")
	router.HandleFunc("/books", controller.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/{bookID}", controller.GetBooksById).Methods("GET")
}
