package handlers

import (
	"fmt"
	"log"
	"net/http"
	"web3/models"
	"web3/pkg/config"
	"web3/pkg/dbdriver"
	"web3/pkg/forms"
	"web3/pkg/repository"
	"web3/pkg/repository/dbrepo"

	render "web3/pkg/render"
)

type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

var Repo *Repository

func NewRepo(appConfig *config.AppConfig, db *dbdriver.DB) *Repository {
	return &Repository{
		App: appConfig,
		DB:  dbrepo.NewPostgresRepo(db.SQL, appConfig),
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) HomeHandler(w http.ResponseWriter, r *http.Request) {
	// id, uid, title, content, _ := m.DB.GetAnArticle()
	// fmt.Println("id :", id)
	// fmt.Println("user id :", uid)
	// fmt.Println("title :", title)
	// fmt.Println("content :", content)
	var artList models.ArticleList
	artList ,err := m.DB.GetThreeArticles()
	if err!=nil{
		log.Println(err)
		return
	}
	for i := range artList.Content{
		fmt.Println(artList.Content[i])
	}
	data := make(map[string]interface{})
	data["articleList"] = artList
	render.RenderTemplate(w, r, "home.page.tmpl", &models.PageData{Data: data})
	m.App.Session.Put(r.Context(), "userid", "shafey")
}

func (m *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	strMap["title"] = "About Us"
	strMap["intro"] = "This page is where we talk about ourselves , we love talking about ourselves"
	strMap["userid"] = m.App.Session.GetString(r.Context(), "userid")
	fmt.Println(m.App.Session.GetString(r.Context(), "userid"))
	render.RenderTemplate(w, r, "about.page.tmpl", &models.PageData{StrMap: strMap})
}
func (m *Repository) LoginHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "login.page.tmpl", &models.PageData{})
}
func (m *Repository) PostLoginHandler(w http.ResponseWriter, r *http.Request) {
	_ = m.App.Session.RenewToken(r.Context())
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	form := forms.New(r.PostForm)
	form.HasRequired("email", "password")
	if !form.Valid() {
		render.RenderTemplate(w, r, "/login", &models.PageData{Form: form})
		return
	}
	id, _, err := m.DB.AuthenicateUser(email, password)
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "invalid email or password")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	m.App.Session.Put(r.Context(), "user_id", id)
	m.App.Session.Put(r.Context(), "flash", "valid login")
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
func (m *Repository) MakePostHandler(w http.ResponseWriter, r *http.Request) {
	if !m.App.Session.Exists(r.Context(), "user_id") {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	var EmptyArticle models.Article
	data := make(map[string]interface{})
	data["article"] = EmptyArticle
	render.RenderTemplate(w, r, "make-post.page.tmpl", &models.PageData{Form: forms.New(nil), Data: data})
}
func (m *Repository) PostMakePostHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	article := models.Post{
		Title:   r.Form.Get("blog_title"),
		Content: r.Form.Get("blog_article"),
		UserId:  2,
	}
	form := forms.New(r.PostForm)
	form.HasRequired("blog_title", "blog_article")
	form.MinLength("blog_title", 5, r)
	form.MinLength("blog_article", 5, r)
	if !form.Valid() {
		data := make(map[string]interface{})
		data["article"] = article
		render.RenderTemplate(w, r, "make-post.page.tmpl", &models.PageData{Form: form, Data: data})
		return
	}
	err = m.DB.InsertPost(article)
	if err != nil {
		log.Fatal(err)
	}
	m.App.Session.Put(r.Context(), "article", article)
	http.Redirect(w, r, "/article-received", http.StatusSeeOther)
}

func (m *Repository) ArticleReceivedHandler(w http.ResponseWriter, r *http.Request) {
	article, ok := m.App.Session.Get(r.Context(), "article").(models.Article)
	if !ok {
		log.Println("Can't get data from session")
		m.App.Session.Put(r.Context(), "error", "can't get data from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	data := make(map[string]interface{})
	data["article"] = article
	render.RenderTemplate(w, r, "article-received.page.tmpl", &models.PageData{Data: data})

}

func (m *Repository) PageHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "page.page.tmpl", &models.PageData{})
}

func (m *Repository) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	_ = m.App.Session.Destroy(r.Context())
	_ = m.App.Session.RenewToken(r.Context())
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
