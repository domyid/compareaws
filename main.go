package main

import (
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	route "github.com/domyid/domyapi/route"
)

func main() {
	http.HandleFunc("/", route.URL)
	adapter := httpadapter.New(http.DefaultServeMux)
	log.Println("Starting Lambda with HTTP adapter...")
	lambda.Start(adapter.Proxy)
}
