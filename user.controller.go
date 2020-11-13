package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
)

//UserHandler function to Handle /users request
func UserHandler(env *Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		userModel := &UserModel{db: env.db}
		switch r.Method {
		case "GET":
			keys, ok := r.URL.Query()["user_id"]
			if !ok || len(keys[0]) < 1 {
				users := userModel.getAllUsers()
				json.NewEncoder(w).Encode(users)
				return
			}
			userID := keys[0]
			user, err := userModel.getUserByID(userID)
			if err == sql.ErrNoRows {
				http.NotFound(w, r)
				return
			}
			json.NewEncoder(w).Encode(user)
			return
		case "POST":
			var newUser User
			json.NewDecoder(r.Body).Decode(&newUser)
			newUser.UserID = uuid.New().String()
			newUser.RegisteredAt = time.Now().Format("2006-01-02 15:04:05")
			if password, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost); err == nil {
				newUser.Password = string(password)
			}
			if err := userModel.insertUser(newUser); err != nil {
				panic(err)
			}
			fmt.Fprint(w, "Sukses Menambahkan User Baru")
			return
		case "DELETE":
			keys, ok := r.URL.Query()["user_id"]
			if !ok || len(keys[0]) < 1 {
				http.Error(w, http.StatusText(404), 404)
				return
			}
			userID := keys[0]
			err := userModel.deleteUserByID(userID)
			if err == sql.ErrNoRows {
				http.NotFound(w, r)
				return
			}
			fmt.Fprint(w, "Data berhasil dihapus")
		}
	})
}
