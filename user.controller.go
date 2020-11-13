package main

import (
	"fmt"
	"net/http"
)

//UserHandler function to Handle /users request
func UserHandler(env *Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userModel := &UserModel{db: env.db}
		switch r.Method {
		case "GET":
			users := userModel.getUserById("1")
			fmt.Println(users)
		}
	})
}
