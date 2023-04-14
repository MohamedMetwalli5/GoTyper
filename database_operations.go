package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host           = "localhost"
	port           = 5432
	user           = "postgres"
	admin_password = "1234"
	dbname         = "GoTyper-DB"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func updateDataBase(username string, password string, wpm string, acc string, raw string) {
	psqlconn := fmt.Sprintf("host= %s port= %d user= %s password= %s dbname= %s sslmode=disable", host, port, user, admin_password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	updateStatement := `UPDATE users
						SET username = $1, password = $2, wpm = $3, acc = $4, raw = $5 
						WHERE username = $1;
						`
	_, e := db.Exec(updateStatement, username, password, wpm, acc, raw)
	CheckError(e)
}

func Login(username string, password string) bool {
	psqlconn := fmt.Sprintf("host= %s port= %d user= %s password= %s dbname= %s sslmode=disable", host, port, user, admin_password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	checkStatement := `SELECT username, password, wpm, acc, raw
					   FROM users
					   where username = $1 and password = $2;
					  `
	row := db.QueryRow(checkStatement, username, password)
	switch err := row.Scan(&username, &password); err {
	case sql.ErrNoRows:
		fmt.Println("Wrong username or password. Try again!")
		return false
	case nil:
		return true
	default:
		fmt.Println("Welcome back,", username+"!")
		return true
	}
}

func SignUp(username string, password string) {
	psqlconn := fmt.Sprintf("host= %s port= %d user= %s password= %s dbname= %s sslmode=disable", host, port, user, admin_password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	insertDynamicStmt := `insert into "users" ("username", "password", "wpm", "acc", "raw") values ($1, $2, $3, $4, $5)`
	_, e := db.Exec(insertDynamicStmt, username, password, 0, 0, 0)
	CheckError(e)
}
