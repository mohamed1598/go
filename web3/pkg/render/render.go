package render

import (
	"fmt"
	"html/template"
	"net/http"
	"web3/models"
	"web3/pkg/config"

	"github.com/justinas/nosurf"
)

var app *config.AppConfig

var tmplCache = make(map[string]*template.Template)

func NewAppConfig(a *config.AppConfig) {
	app = a
}

func AddCSRFData(pd *models.PageData, r *http.Request) *models.PageData {
	pd.CSRFToken = nosurf.Token(r)
	if app.Session.Exists(r.Context(), "user_id") {
		pd.IsAuthenticated = 1
	}
	return pd
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, t string, pageData *models.PageData) {
	var tmpl *template.Template
	var err error
	_, inMap := tmplCache[t]
	if !inMap {
		err = makeTemplate(t)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("template in cache")
		}
	}
	tmpl = tmplCache[t]
	pageData = AddCSRFData(pageData, r)
	fmt.Println(pageData)
	err = tmpl.Execute(w, pageData)
	if err != nil {
		fmt.Println(err)
	}
}

func makeTemplate(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"templates/base.layout.tmpl",
	}
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	tmplCache[t] = tmpl
	return nil
}
