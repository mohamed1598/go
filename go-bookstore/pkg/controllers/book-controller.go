package controllers

import (
	"bookstore/pkg/models"
	"bookstore/pkg/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("inside book route")
	newBooks, err := models.GetAllBooks()
	// fmt.Println("inside get all books",newBooks,err)
	if err != nil {
		log.Fatal("there is a problem in getting all book")
	}
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetBookBYID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(int(Id))
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	createBook, err := models.CreateBook(*CreateBook)
	if err != nil {
		fmt.Println(err)
	}
	res, _ := json.Marshal(createBook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	err = models.DeleteBook(int(Id))
	if err != nil {
		fmt.Println(err)
	}
	res, _ := json.Marshal("book deleter successfully")
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter,r *http.Request){
	var UpdateBook = models.Book{}
	utils.ParseBody(r,&UpdateBook)
	UpdateBook,err := models.UpdateBook(UpdateBook)
	if err != nil {
		fmt.Println(err)
	}
	res, _ := json.Marshal(UpdateBook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}