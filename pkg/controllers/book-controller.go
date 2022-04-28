package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	book := models.DeleteBookById(ID)
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(book)
	w.Write(res)

}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["id"]
	id,err := strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Println(err)
	}
	bookDetails, _ := models.GetBookById(id)
	w.WriteHeader(http.StatusOK)
	book,_:=json.Marshal(bookDetails)
	w.Write(book)
}
