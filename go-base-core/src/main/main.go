package main

import (
	"net/http"

	"github.com/gorilla/context"
)

func main() {
	configRoute()
	http.ListenAndServe(":6868", context.ClearHandler(http.DefaultServeMux))
}
