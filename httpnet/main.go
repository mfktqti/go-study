package main

import (
	"net/http"

	"github.com/gorilla/context"
)

func main() {
	initRoute()
	http.ListenAndServe(":6868", context.ClearHandler(http.DefaultServeMux))
}

func getResponse(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.RequestURI))
}

func initRoute() {
	http.Handle("/m", http.HandlerFunc(getResponse))
	http.Handle("/n/", http.StripPrefix("/n/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})))

	//http.Handle("/n/", http.StripPrefix("/n/", htt))
}

// AllowsPostOnly is an http middleware that allows http request with GET method only
func AllowsPostOnly(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			http.Error(w, "Method not allowed", 405)
			return
		}
		next.ServeHTTP(w, r)
	})
}
