package db

import (
	"fmt"
	"log"

	"github.com/mohamed1598/golang-gorm-postgres/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func Init() *gorm.DB{
	var err error
	// const connectionString = "host=localhost port=5432 dbname=blog_db user=postgres password=123456"
	dsn := "host=localhost port=5432 dbname=gorm_db user=postgres password=123456"
	// gorm.Dialector
	db,err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err != nil{
		log.Fatal("failed to connect to database error :",err)
	}
	fmt.Println("? connected successfully to the database")
	db.AutoMigrate(&models.Book{})
	return db
} 