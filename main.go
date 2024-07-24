package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	route "github.com/domyid/domyapi/route"
)

type Handler struct{}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	route.URL(w, r)
}

func main() {
	handle := new(Handler)
	http.Handle("/", handle)
	adapter := httpadapter.New(http.DefaultServeMux)
	lambda.Start(adapter.Proxy)
}
