package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"time"
	"web3/models"
	"web3/pkg/config"
	"web3/pkg/dbdriver"
	"web3/pkg/handlers"
	"web3/pkg/helpers"
	"web3/pkg/render"

	"github.com/alexedwards/scs/v2"
)

var sessionManager *scs.SessionManager

const connectionString = "host=localhost port=5432 dbname=blog_db user=postgres password=123456"

func main() {
	// pw := "123456"
	// hpw, err := bcrypt.GenerateFromPassword([]byte(pw), 12)
	// fmt.Println(string(hpw))
	var app config.AppConfig
	db, err := run(&app)
	defer db.SQL.Close()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	helpers.ErrorCheck(err)
	// http.HandleFunc("/", handlers.HomeHandler)
	// http.HandleFunc("/about", handlers.AboutHandler)
	// err := http.ListenAndServe("localhost:8080", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func run(app *config.AppConfig) (*dbdriver.DB, error) {

	gob.Register(models.Article{})
	gob.Register(models.User{})
	gob.Register(models.Post{})

	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.Secure = false
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	app.Session = sessionManager

	db, err := dbdriver.ConnectSQL(connectionString)
	if err != nil {
		log.Fatal("can't connect to db")
	}

	repo := handlers.NewRepo(app, db)
	handlers.NewHandlers(repo)

	render.NewAppConfig(app)
	return db, err
}
