package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func getAllRowData(conn *sql.DB) error {
	rows, err := conn.Query("select id , email ,name from users")
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()
	var name, email string
	var id int
	for rows.Next() {
		err = rows.Scan(&id, &email, &name)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println("Data :", id, name, email)
		}
	}
	if err != nil {
		log.Fatal("error reading data", err)
	}
	return nil

}

func getUserData(conn *sql.DB, userId int) error {
	var name, email, pw string
	var id, uType int
	query := fmt.Sprintf(`select id,name,email,password,user_type from users where id = %d`, userId)
	row := conn.QueryRow(query)
	err := row.Scan(&id, &name, &email, &pw, &uType)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("Id :", id)
	fmt.Println("Name :", name)
	fmt.Println("password :", pw)
	fmt.Println("user_type :", uType)
	fmt.Println("email :", email)
	return nil
}

func InsertNewUser(conn *sql.DB, name string, email string, pw string, uType int) error {
	query := fmt.Sprintf(`insert into users(name,email,password,acc_created,last_login,user_type) VALUES ('%s','%s','%s',current_timestamp,current_timestamp,%d)`, name, email, pw, uType)
	_, err := conn.Exec(query)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func UpdateUserName(conn *sql.DB, newName string, id int) error {
	query := fmt.Sprintf(`update users set name='%s' where id=%d`, newName, id)
	_, err := conn.Exec(query)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func DeleteUser(conn *sql.DB, id int) error {
	query := fmt.Sprintf(`delete from users where id=%d`, id)
	_, err := conn.Exec(query)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func main() {
	//connect to database
	conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=blog_db user=postgres password=123456")
	if err != nil {
		log.Fatalf(fmt.Sprintf("couldn't connect to db : %v\n", err))
	}
	defer conn.Close()
	err = conn.Ping()
	if err != nil {
		log.Fatalf("cant connect to db")
	}
	err = getAllRowData(conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-------------------")
	// err = InsertNewUser(conn, "ahmed", "ahmed@mail.com", "123456", 1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = getAllRowData(conn)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// UpdateUserName(conn, "hassan", 2)
	// fmt.Println("-------------------")
	// getUserData(conn, 2)
	// DeleteUser(conn, 1)
}
