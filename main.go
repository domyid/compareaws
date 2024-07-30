package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/domyid/handler"
)

func main() {
	lambda.Start(handler.Login)
}
