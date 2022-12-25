package main

import (
	"bookstore/pkg/config"
	"bookstore/pkg/models"
	"bookstore/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var connectionString = "host=localhost port=5432 dbname=bookstore_db user=postgres password=123456"

func main() {
	db, err := config.ConnectSQL(connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()
	models.IntializeModels(db.SQL)
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("executing program")
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}
