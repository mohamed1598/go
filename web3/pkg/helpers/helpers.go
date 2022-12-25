package helpers

import (
	"log"
	"net/http"
	"web3/pkg/config"
)

var app config.AppConfig

func ErrorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func IsAuthenticated(r *http.Request) bool{
	exists := app.Session.Exists(r.Context(),"user_id")
	return exists
}
