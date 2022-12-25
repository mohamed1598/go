package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"web3/pkg/helpers"

	"github.com/justinas/nosurf"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !helpers.IsAuthenticated(r) {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			log.Fatal("error logging in")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func LogRequestInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		fmt.Printf("%d/%d/%d : %d:%d:%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
		fmt.Println(r.URL.Path)
		fmt.Println("form title", r.Form.Get("blog_title"))
		next.ServeHTTP(w, r)
	})
}
func SetupSession(next http.Handler) http.Handler {
	return sessionManager.LoadAndSave(next)
}

func NoSurf(next http.Handler) http.Handler {
	NoSurfHandler := nosurf.New(next)
	NoSurfHandler.SetBaseCookie(http.Cookie{
		Name:     "myCSRFCookie",
		Path:     "/",
		Domain:   "",
		Secure:   false,
		HttpOnly: true,
		MaxAge:   3600,
		SameSite: http.SameSiteLaxMode,
	})
	return NoSurfHandler
}
