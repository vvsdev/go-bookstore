package main

import (
	"net/http"
	"fmt"
)

func main() {
	routeHandler()
}

func routeHandler() {
	// mux := http.DefaultServeMux

	http.HandleFunc("/", tesRoute)

	addr := "127.0.0.1:5000"
	server := new(http.Server)
	server.Addr = addr

	fmt.Println("Server started at " + addr)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func tesRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HelloWorld"))
}