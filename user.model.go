package main

import (
	"database/sql"
	"fmt"
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

func (model *UserModel) getAllUsers() []User {
	fakeModel := []User{
		User{
			UserID:       "1",
			FirstName:    "A",
			LastName:     "B",
			Age:          13,
			RegisteredAt: "Now",
			Address:      "C",
			Email:        "D",
			Password:     "E",
		},
		User{
			UserID:       "1",
			FirstName:    "A",
			LastName:     "B",
			Age:          13,
			RegisteredAt: "Now",
			Address:      "C",
			Email:        "D",
			Password:     "E",
		},
	}
	return fakeModel
}

func (model *UserModel) getUserById(id string) User {
	query := fmt.Sprintf("select * from users where user_id=%s", id)
	res, _ := model.db.Query(query)
	fmt.Print(res)
	return User{}
}
