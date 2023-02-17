package controller

import (
	"encoding/json"
	"net/http"
	model "sever-client/src/models"
	"sever-client/src/utils"
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
