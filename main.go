package main

import (
	"bytes"
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	route "github.com/domyid/domyapi/route"
)

// LambdaResponseWriter adalah custom ResponseWriter untuk Lambda
type LambdaResponseWriter struct {
	statusCode int
	headers    http.Header
	body       bytes.Buffer
}

func NewLambdaResponseWriter() *LambdaResponseWriter {
	return &LambdaResponseWriter{
		headers: http.Header{},
	}
}

func (r *LambdaResponseWriter) Header() http.Header {
	return r.headers
}

func (r *LambdaResponseWriter) Write(b []byte) (int, error) {
	return r.body.Write(b)
}

func (r *LambdaResponseWriter) WriteHeader(statusCode int) {
	r.statusCode = statusCode
}

func (r *LambdaResponseWriter) GetResponse() events.APIGatewayProxyResponse {
	headers := make(map[string]string)
	for k, v := range r.headers {
		if len(v) > 0 {
			headers[k] = v[0]
		}
	}
	return events.APIGatewayProxyResponse{
		StatusCode: r.statusCode,
		Headers:    headers,
		Body:       r.body.String(),
	}
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	writer := NewLambdaResponseWriter()
	req, err := http.NewRequest(request.HTTPMethod, request.Path, bytes.NewBufferString(request.Body))
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Failed to create request",
		}, nil
	}

	for k, v := range request.Headers {
		req.Header.Set(k, v)
	}

	route.URL(writer, req)

	return writer.GetResponse(), nil
}

func main() {
	lambda.Start(handler)
}
