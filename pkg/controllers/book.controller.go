package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dixi225/go-fiber-postgres/pkg/models"
	"github.com/dixi225/go-fiber-postgres/pkg/utils"
	"github.com/gorilla/mux"
)

var newBook models.Book

func getBook(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBook()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func getBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	book := models.GetBookById(int(ID))
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(&book)
	w.Write(res)
}
