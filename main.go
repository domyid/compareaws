package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	route "github.com/domyid/domyapi/route"
)

type RetMsg struct {
	Message string `json:"message"`
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Buat response writer dan request dari APIGatewayProxyRequest
	writer := NewLambdaResponseWriter()
	req, err := http.NewRequest(request.HTTPMethod, request.Path, nil)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Failed to create request",
		}, nil
	}

	// Panggil fungsi route.URL untuk menangani permintaan
	route.URL(writer, req)

	// Kembalikan response yang dibuat oleh writer
	return writer.GetResponse(), nil
}

func main() {
	lambda.Start(handler)
}
