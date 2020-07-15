package main

import (
	"database/sql"
	"fmt"
)

func seedDB(db *sql.DB) {
	query := []string{
		`insert into publishers values(
			'1', 'UB PRESS'
		);`,
		`insert into authors values(
			'1', 'Ucok Lewandowski'
		);`,
		`insert into categories values(
			'1', 'Ensiklopedia', '0'
		);`,
		`insert into books values(
			'1', 'Sejarah Pohon Pisang', 'Pohon pisang dulunya adalah pohon Nangka', '1', '1', 200000, 1
		);`,
		`insert into book_authors values(
			'1', '1'
		);`,
		`insert into users values(
			'1', 'Budi', 'Stanford', 20, '1999-09-09', 'Indonesia', 'budistanford@gmail.com', 'YnVkaXN0YW5mb3JkcGVjaW50YWt1Y2luZw=='
		);`,
		`insert into transactions values(
			'1', '1', '2020-02-20'
		);
		`}
	for _, q := range query {
		_, err := db.Exec(q)
		if err != nil {
			panic(err.Error())
		}
	}

	fmt.Println("Seeding Success")
}

func unseedDB(db *sql.DB) {
	query := []string{
		`delete from publishers;`,
		`delete from authors;`,
		`delete from categories;`,
		`delete from books;`,
		`delete from book_authors;`,
		`delete from users;`,
		`delete from transactions;
		`}
	for _, q := range query {
		_, err := db.Exec(q)
		if err != nil {
			panic(err.Error())
		}
	}

	fmt.Println("Unseeding Success")
}
