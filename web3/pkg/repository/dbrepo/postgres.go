package dbrepo

import (
	"context"
	"errors"
	"time"
	"web3/models"

	"golang.org/x/crypto/bcrypt"
)

func (m *postgresDbRepostory) InsertPost(newPost models.Post) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := `insert into posts (title,content,user_id) values ($1,$2,$3)`
	_, err := m.DB.ExecContext(ctx, query, newPost.Title, newPost.Content, newPost.UserId)
	return err
}

func (m *postgresDbRepostory) GetAnArticle() (int, int, string, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var id, uId int
	var aTitle, aContent string
	query := `select id,user_id,title,content from posts limit 1`
	row := m.DB.QueryRowContext(ctx, query)
	err := row.Scan(&id, &uId, &aTitle, &aContent)
	if err != nil {
		return id, uId, "", "", err
	}
	return id, uId, aTitle, aContent, nil
}

func (m *postgresDbRepostory) GetThreeArticles() (models.ArticleList, error) {
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var artList models.ArticleList
	rows, err := m.DB.Query("select id,user_id,title,content from posts order by id Desc Limit $1", 3)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id, uId int
		var title, content string
		err = rows.Scan(&id, &uId, &title, &content)
		if err != nil {
			panic(err)
		}
		artList.Id = append(artList.Id, id)
		artList.UserId = append(artList.UserId, uId)
		artList.Title = append(artList.Title, title)
		artList.Content = append(artList.Content, content)
	}
	err = rows.Err()
	return artList, err
}

func (m *postgresDbRepostory) GetUserById(id int) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := `select name,email,password,acct_created,last_login,user_type,id from usere where id =$1`
	row := m.DB.QueryRowContext(ctx, query, id)
	var u models.User
	err := row.Scan(&u.Name, &u.Email, &u.Password, &u.AcctCreated, &u.LastLogin, &u.UserType, &u.Id)
	return u, err
}
func (m *postgresDbRepostory) UpdateUser(u models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := `update users set name=$1,email=$2,last_login=$3,user_type=$4,id=$5 where id=$5`
	_, err := m.DB.ExecContext(ctx, query, u.Name, u.Email, u.LastLogin, u.UserType, u.Id)
	return err
}
func (m *postgresDbRepostory) AuthenicateUser(testEmail, testPassword string) (int, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var id int
	var hashedPw string
	query := `select id,password from users where email=$1`
	row := m.DB.QueryRowContext(ctx, query, testEmail)
	err := row.Scan(&id, &hashedPw)
	if err != nil {
		return id, "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPw), []byte(testPassword))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, "", errors.New("password is not correct")
	} else if err != nil {
		return 0, "", err
	}
	return id, hashedPw, nil
}
