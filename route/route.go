package domyApi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	configs "github.com/domyid/domyapi/config"
)

var lambdaClient *lambda.Client

func init() {
	// Initialize AWS Lambda client
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		fmt.Printf("unable to load SDK config, %v", err)
	}

	lambdaClient = lambda.NewFromConfig(cfg)
}

func invokeLambda(functionName string, payload []byte) ([]byte, error) {
	input := &lambda.InvokeInput{
		FunctionName: aws.String(functionName),
		Payload:      payload,
	}

	result, err := lambdaClient.Invoke(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to invoke function %s, %w", functionName, err)
	}

	if result.FunctionError != nil {
		return nil, fmt.Errorf("function %s returned an error: %s", functionName, aws.ToString(result.FunctionError))
	}

	return result.Payload, nil
}

func URL(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if configs.SetAccessControlHeadersForLambda(&request) {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNoContent,
			Headers:    request.Headers,
		}, nil // If it's a preflight request, return early.
	}

	var method, path string = request.HTTPMethod, request.Path

	var response events.APIGatewayProxyResponse

	switch {
	case method == "POST" && path == "/login":
		response = invokeLambdaHandler(request, "loginLambdaFunction")
	case method == "POST" && path == "/refresh-token":
		response = invokeLambdaHandler(request, "refreshTokenLambdaFunction")
	case method == "GET" && path == "/data/mahasiswa":
		response = invokeLambdaHandler(request, "getMahasiswaLambdaFunction")
	case method == "GET" && path == "/data/bimbingan/mahasiswa":
		response = invokeLambdaHandler(request, "getListBimbinganMahasiswaLambdaFunction")
	case method == "POST" && path == "/data/bimbingan/mahasiswa":
		response = invokeLambdaHandler(request, "postBimbinganMahasiswaLambdaFunction")
	case method == "GET" && path == "/data/dosen":
		response = invokeLambdaHandler(request, "getDosenLambdaFunction")
	case method == "POST" && path == "/jadwalmengajar":
		response = invokeLambdaHandler(request, "getJadwalMengajarLambdaFunction")
	case method == "POST" && path == "/riwayatmengajar":
		response = invokeLambdaHandler(request, "getRiwayatPerkuliahanLambdaFunction")
	case method == "POST" && path == "/absensi":
		response = invokeLambdaHandler(request, "getAbsensiKelasLambdaFunction")
	case method == "POST" && path == "/nilai":
		response = invokeLambdaHandler(request, "getNilaiMahasiswaLambdaFunction")
	case method == "POST" && path == "/BAP":
		response = invokeLambdaHandler(request, "getBAPLambdaFunction")
	case method == "GET" && path == "/data/list/ta":
		response = invokeLambdaHandler(request, "getListTugasAkhirMahasiswaLambdaFunction")
	case method == "POST" && path == "/data/list/bimbingan":
		response = invokeLambdaHandler(request, "getListBimbinganMahasiswaLambdaFunction")
	case method == "POST" && path == "/approve/bimbingan":
		response = invokeLambdaHandler(request, "approveBimbinganLambdaFunction")
	default:
		response = events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
			Body:       "Not Found",
		}
	}
	return response, nil
}

func invokeLambdaHandler(request events.APIGatewayProxyRequest, functionName string) events.APIGatewayProxyResponse {
	payload, err := json.Marshal(request.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf("Failed to marshal request body: %v", err),
		}
	}

	responsePayload, err := invokeLambda(functionName, payload)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf("Failed to invoke lambda: %v", err),
		}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(responsePayload),
		Headers:    request.Headers,
	}
}
