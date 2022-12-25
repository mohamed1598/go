package dbdriver

import (
	"database/sql"
	"time"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}
const maxOpenDbConns = 20
const maxIdleConns = 10
const maxDbLifeTime = 5 * time.Minute

func NewDatabase(dsn string)(*sql.DB,error){
	db,err := sql.Open("pgx",dsn)
	if err != nil{
		return nil,err
	}
	err = db.Ping()
	if err != nil{
		return nil,err
	}
	return db,nil
}
func ConnectSQL(dsn string)(*DB,error){
	db,err:= NewDatabase(dsn)
	if err != nil{
		panic(err)
	}
	db.SetMaxOpenConns(maxOpenDbConns)
	db.SetConnMaxIdleTime(maxIdleConns)
	db.SetConnMaxLifetime(maxDbLifeTime)
	dbConn.SQL = db
	return dbConn,nil
}