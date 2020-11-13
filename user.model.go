package main

import (
	"database/sql"
	"log"
)

//UserModel Constructor
type UserModel struct {
	db *sql.DB
}

//User model
type User struct {
	UserID       string `json:"id"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Age          int8   `json:"age"`
	RegisteredAt string `json:"registeredat"`
	Address      string `json:"address"`
	Email        string `json:"email"`
	Password     string `json:"password"`
}

func (model *UserModel) getAllUsers() []*User {
	users := make([]*User, 0)

	rows, err := model.db.Query("SELECT * FROM users")

	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		user := new(User)
		err := rows.Scan(
			&user.UserID,
			&user.FirstName,
			&user.LastName,
			&user.Age,
			&user.RegisteredAt,
			&user.Address,
			&user.Email,
			&user.Password,
		)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	return users
}

func (model *UserModel) getUserByID(id string) (*User, error) {
	row := model.db.QueryRow("select * from users where user_id = ?", id)
	user := new(User)
	err := row.Scan(
		&user.UserID,
		&user.FirstName,
		&user.LastName,
		&user.Age,
		&user.RegisteredAt,
		&user.Address,
		&user.Email,
		&user.Password)
	return user, err
}

func (model *UserModel) insertUser(user User) error {
	_, err := model.db.Exec("INSERT INTO users VALUES(?, ?, ?, ?, ?, ?, ?, ?)",
		user.UserID,
		user.FirstName,
		user.LastName,
		user.Age,
		user.RegisteredAt,
		user.Address,
		user.Email,
		user.Password)
	return err
}

func (model *UserModel) deleteUserByID(id string) error {
	_, err := model.db.Exec("delete from users where user_id = ?", id)
	return err
}
