package main

import (
	"net/http"
	"web3/pkg/config"
	"web3/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) *chi.Mux {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(LogRequestInfo)
	mux.Use(NoSurf)
	mux.Use(SetupSession)
	mux.Get("/", handlers.Repo.HomeHandler)
	mux.Get("/about", handlers.Repo.AboutHandler)
	mux.Get("/login", handlers.Repo.LoginHandler)
	mux.Post("/login", handlers.Repo.PostLoginHandler)
	mux.Get("/make-post", handlers.Repo.MakePostHandler)
	mux.Post("/make-post", handlers.Repo.PostMakePostHandler)
	mux.Get("/page", handlers.Repo.PageHandler)
	mux.Get("/article-received", handlers.Repo.ArticleReceivedHandler)
	mux.Get("/logout", handlers.Repo.LogoutHandler)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
