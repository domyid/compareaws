package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	route "github.com/domyid/domyapi/route"
)

func main() {
	http.HandleFunc("WebHook", route.URL)
	adapter := httpadapter.New(http.DefaultServeMux)
	lambda.Start(adapter.Proxy)
}
