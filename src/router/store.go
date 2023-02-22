package router

import (
	"sever-client/src/controller"

	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controller.CreateBook).Methods("POST")
	router.HandleFunc("/books", controller.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/{bookID}", controller.GetBooksById).Methods("GET")
	router.HandleFunc("/books/{bookID}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/del/{bookID}", controller.DeleteBook).Methods("DELETE")
}

//line 12 and 13 has same paths, this might cause a "405 method not allowed" error
