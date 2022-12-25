package dbrepo

import (
	"database/sql"
	"web3/pkg/config"
	"web3/pkg/repository"
)

type postgresDbRepostory struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, ac *config.AppConfig) repository.DatabaseRepo {
	return &postgresDbRepostory{
		App: ac,
		DB:  conn,
	}
}
