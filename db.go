package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func getDB() *sql.DB {
	err := godotenv.Load()
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbInit := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
	db, err := sql.Open("mysql", dbInit)
	err = db.Ping()

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("DB Connection Established")
	return db
}

func migrateDB(db *sql.DB) {
	query := []string{
		`create table if not exists publishers(
			publisher_id varchar(100) not null primary key,
			publisher varchar(50)
		);`,
		`create table if not exists authors(
			author_id varchar(100) not null primary key,
			author varchar(50)
		);`,
		`create table if not exists categories(
			category_id varchar(100) not null primary key,
			category varchar(20),
			is_default char(1)
		);`,
		`create table if not exists books(
			book_id varchar(100) not null primary key,
			title varchar(255),
			synopsis varchar(500),
			publisher_id varchar(100) references publishers(publisher_id),
			category_id varchar(100) references categories(category_id),
			price long,
			stock long
		)
		`,
		`create table if not exists book_authors(
			author_id varchar(100) references authors(author_id),
			book_id varchar(100) references books(book_id),
			primary key(author_id, book_id)
		)
		`,
		`create table if not exists users(
			user_id varchar(100) not null primary key,
			first_name varchar(50),
			last_name varchar(50),
			age smallint,
			registeredAt timestamp,
			address varchar(500),
			email varchar(50),
			password varchar(100)
		)
		`,
		`create table if not exists transactions(
			user_id varchar(100) references users(user_id),
			book_id varchar(100) references books(book_id),
			time timestamp,
			primary key(user_id, book_id, time)
		)
		`}
	for _, q := range query {
		_, err := db.Exec(q)
		if err != nil {
			panic(err.Error())
		}
	}

	fmt.Println("Migration Success")
}
