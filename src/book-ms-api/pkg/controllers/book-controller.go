package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/philipokiokio/book-ms-api/pkg/models"
	"github.com/philipokiokio/book-ms-api/pkg/utils"
)

var NewBook models.Book

func getBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()

	res, _ := json.Marshal((newBooks))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func getBookById(w http.ResponseWriter, r *http.Request) {

	vars := mux.vars(r)
	bookId := vars["bookId"]

	Id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	book, _ := models.GetBook(Id)

	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func deleteBookById(w http.ResponseWriter, r *http.Request) {

	vars := mux.vars(r)
	bookId := vars["bookId"]

	Id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	models.DeleteBook(Id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}

	utils.ParseBody(r, createBook)

	b := createBook.CreateBook()

	res, _ := json.Marshal(b)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func updateBook(w http.ResponseWriter, r *http.Request) {

	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)

	vars := mux.vars(r)
	bookId := vars["bookId"]

	Id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, db := models.GetBook(Id)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author

	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication

	}
	db.Save(&bookDetails)

	res, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
