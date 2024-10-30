package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

// "docker run --name my-postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=1234 -p 5432:5432 -d postgres"

var (
	host           string
	port           int
	user           string
	admin_password string
	dbname         string
)

func InitiateDataBaseVariables() {
	host = os.Getenv("host")
	port, _ = strconv.Atoi(os.Getenv("port"))
	user = os.Getenv("user")
	admin_password = os.Getenv("admin_password")
	dbname = os.Getenv("dbname")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func createDatabase() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, admin_password)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	// Attempt to create the database
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE \"%s\"", dbname))
}

func createUsersTable() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, admin_password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	// Create the users table if it doesn't exist
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS public.users (
		username character varying(255) NOT NULL,
		password character varying(255),
		wpm integer DEFAULT 0,
		acc integer DEFAULT 0,
		"raw" integer DEFAULT 0,
		CONSTRAINT users_pkey PRIMARY KEY (username)
	);`
	_, err = db.Exec(createTableSQL)
	CheckError(err)
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
