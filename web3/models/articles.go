package models

type Article struct {
	Id      int
	UserId  int
	Title   string
	Content string
}

type ArticleList struct {
	Id      []int
	UserId  []int
	Title   []string
	Content []string
}
