package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"
)

//Env for project global environment
type Env struct {
	db *sql.DB
}

func main() {
	db := getDB()
	env := &Env{db: db}
	seed := flag.Bool("seed", false, "seeding DB")
	unseed := flag.Bool("unseed", false, "unseeding DB")
	migrate := flag.Bool("migrate", false, "migrate DB")

	flag.Parse()

	if *migrate {
		migrateDB(db)
		return
	}
	if *seed {
		seedDB(db)
		return

	} else if *unseed {
		unseedDB(db)
		return
	}

	//Router definition
	//Auth
	http.Handle("/auth/login", login(env))

	//User
	http.Handle("/users", UserHandler(env))

	addr := "127.0.0.1:5000"
	server := new(http.Server)
	server.Addr = addr

	fmt.Println("Server started at " + addr)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()
}
