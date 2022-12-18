package main

import (
	"net/http"
)

type UserRoute struct{}

func (u *UserRoute) Init() {
	http.Handle("/user", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world ! user"))
	}))
}
