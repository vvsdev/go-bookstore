package main

import (
	"database/sql"
	"fmt"
)

func seedDB(db *sql.DB) {
	querys := []string{
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
			'44968a50-957f-40c9-a949-20a8ac486556', 'Ibrahimsyah', 'zairussalam', 20, '2020-11-13 15:31:21', 'malang', 'ibrahim@gmail.com', '$2a$10$erCcnN2PIx9FqqA7IHFsa.BMN2.ASrUEB2.v6N7WKWSsLMtNWXmyi'
		),
		(
			'ae115d4f-da81-4180-9976-c0b04726cbbd', 'Devian Wahyu', 'Setiyawan', 20, '2020-11-13 15:32:16', 'tulungagung', 'devian@wahyu.setiawan', '$2a$10$shpr1u7g8H/mgf.o0dP3D.SU98UWRRsblUmriB9BRH9kuw5Vcv/ES'
		);`,
		`insert into transactions values(
			'1', '1', '2020-02-20'
		);
		`}
	for _, query := range querys {
		_, err := db.Exec(query)
		if err != nil {
			panic(err.Error())
		}
	}

	fmt.Println("Seeding Success")
}

func unseedDB(db *sql.DB) {
	querys := []string{
		`delete from publishers;`,
		`delete from authors;`,
		`delete from categories;`,
		`delete from books;`,
		`delete from book_authors;`,
		`delete from users;`,
		`delete from transactions;
		`}
	for _, query := range querys {
		_, err := db.Exec(query)
		if err != nil {
			panic(err.Error())
		}
	}

	fmt.Println("Unseeding Success")
}
