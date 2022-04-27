package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/rabigautam/book_management_system/pkg/models"
	"github.com/rabigautam/book_management_system/pkg/utils"
)

var NewBook models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
