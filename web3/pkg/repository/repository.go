package repository

import "web3/models"

type DatabaseRepo interface {
	InsertPost(newPost models.Post) error
	GetAnArticle() (int, int, string, string, error)
	GetThreeArticles() (models.ArticleList, error)
	GetUserById(id int) (models.User, error)
	UpdateUser(u models.User) error
	AuthenicateUser(testEmail, testPassword string) (int, string, error)
}
