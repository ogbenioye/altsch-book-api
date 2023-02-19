package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	model "sever-client/src/models"
	"sever-client/src/utils"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBook model.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {
	Book := &model.Book{}
	utils.ParseBody(r, Book)
	b := Book.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := model.GetAllBooks()
	resp, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func GetBooksById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	bookId := vars["bookID"]

	Id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing")
	}

	book := model.GetBooksById(int(Id))

	resp, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
