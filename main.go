package main

import (
	"net/http"

	"github.com/akrylysov/algnhsa"
	route "github.com/domyid/domyapi/route"
)

type Handler struct{}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	route.URL(w, r)
}

func main() {
	handle := new(Handler)
	http.Handle("/", handle)
	algnhsa.ListenAndServe(http.DefaultServeMux, nil)
}
