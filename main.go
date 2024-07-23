package domyApi

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	route "github.com/domyid/domyapi/route"
)

func WebHookHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    // Convert APIGatewayProxyRequest to http.Request
    req, err := http.NewRequest(request.HTTPMethod, request.Path, nil)
    if err != nil {
        return events.APIGatewayProxyResponse{}, err
    }
    // Call the actual handler
    rr := &responseRecorder{}
    route.URL(rr, req)
    
    // Convert http.Response to APIGatewayProxyResponse
    return events.APIGatewayProxyResponse{
        StatusCode: rr.status,
        Headers:    map[string]string{"Content-Type": "application/json"},
        Body:       string(rr.body),
    }, nil
}

func main() {
    lambda.Start(WebHookHandler)
}

// responseRecorder is an implementation of http.ResponseWriter
type responseRecorder struct {
    status int
    body   []byte
}

func (r *responseRecorder) Header() http.Header {
    return http.Header{}
}

func (r *responseRecorder) Write(b []byte) (int, error) {
    r.body = append(r.body, b...)
    return len(b), nil
}

func (r *responseRecorder) WriteHeader(statusCode int) {
    r.status = statusCode
}
