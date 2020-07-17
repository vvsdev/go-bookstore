package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	db := getDB()
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

	http.HandleFunc("/", mainRouter)
	fs := http.FileServer(http.Dir("./swagger"))
	http.Handle("/swagger/", http.StripPrefix("/swagger/", fs))
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
