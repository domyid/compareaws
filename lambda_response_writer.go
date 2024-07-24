package main

import (
	"bytes"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

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
