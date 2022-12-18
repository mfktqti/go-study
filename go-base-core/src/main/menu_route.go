package main

import (
	"fmt"
	"net/http"
)

type MenuRoute struct {
}

func (m *MenuRoute) Init() {
	fmt.Println("MenuRoute Init()")
	http.Handle("/menu", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world ! menu！"))
	}))
}
