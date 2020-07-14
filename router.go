package main

import "net/http"

func mainRouter(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
