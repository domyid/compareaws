package main

import (
	"log"
	"net/http"

	route "github.com/domyid/domyapi/route"
)

func main() {
	http.HandleFunc("/", route.URL)
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
