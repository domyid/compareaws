package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	route "github.com/domyid/domyapi/route"
)

func main() {
	lambda.Start(route.URL)
}
