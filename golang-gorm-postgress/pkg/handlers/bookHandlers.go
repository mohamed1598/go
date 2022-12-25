package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mohamed1598/golang-gorm-postgres/models"
)

func (h Handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	if result := h.DB.Find(&books); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}
func (h Handler) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var book models.Book
	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
func (h Handler) AddBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var book models.Book
	json.Unmarshal(body, &book)
	if result := h.DB.Create(&book); result.Error != nil {
		fmt.Println(result.Error)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("created")
}
func (h Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var updatedBook, book models.Book
	json.Unmarshal(body, &updatedBook)
	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}
	book.Title = updatedBook.Title
	book.Author = updatedBook.Author
	book.Desc = updatedBook.Desc
	h.DB.Save(&book)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}
func (h Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var book models.Book
	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}
	h.DB.Delete(&book)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}
