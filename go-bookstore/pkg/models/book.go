package models

import (
	"context"
	"database/sql"
	"time"
)

var db *sql.DB

type Book struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func IntializeModels(database *sql.DB) {
	db = database
}
func CreateBook(book Book) (Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := `insert into books(name,author,publication) values ($1,$2,$3)`
	_, err := db.ExecContext(ctx, query, book.Name, book.Author, book.Publication)
	return book, err
}
func GetAllBooks() ([]Book, error) {
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var books []Book
	rows, err := db.Query("select id,name,author,publication from books")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var book Book
		err = rows.Scan(&book.Id, &book.Name, &book.Author, &book.Publication)
		if err != nil {
			panic(err)
		}
		books = append(books, book)
	}
	err = rows.Err()
	return books, err

}
func GetBookById(id int) (Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := `select id,name,author,publication from books where id=$1`
	row := db.QueryRowContext(ctx, query, id)
	var book Book
	err := row.Scan(&book.Id, &book.Name, &book.Author, &book.Publication)
	return book, err
}
func DeleteBook(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := `delete from books where id =$1`
	_, err := db.ExecContext(ctx, query, id)
	return err
}
func UpdateBook(book Book) (Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := `update books set name=$1,author=$2,publication=$3 where id =$4`
	_, err := db.ExecContext(ctx, query, book.Name, book.Author, book.Publication, book.Id)
	return book, err
}
