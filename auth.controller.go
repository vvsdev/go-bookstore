package main

import (
	"net/http"
)

func login(env *Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Login"))
	})
}
