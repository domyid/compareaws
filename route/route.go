package domyApi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	configs "github.com/domyid/domyapi/config"
)

var lambdaClient *lambda.Client

func init() {
	// Initialize AWS Lambda client
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-southeast-2"))
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

func URL(w http.ResponseWriter, r *http.Request) {
	if configs.SetAccessControlHeaders(w, r) {
		return // If it's a preflight request, return early.
	}

	var method, path string = r.Method, r.URL.Path

	switch {
	case method == "POST" && path == "/login":
		invokeLambdaHandler(w, r, "loginLambdaFunction")
	case method == "POST" && path == "/refresh-token":
		invokeLambdaHandler(w, r, "refreshTokenLambdaFunction")
	case method == "GET" && path == "/data/mahasiswa":
		invokeLambdaHandler(w, r, "getMahasiswaLambdaFunction")
	case method == "GET" && path == "/data/bimbingan/mahasiswa":
		invokeLambdaHandler(w, r, "getListBimbinganMahasiswaLambdaFunction")
	case method == "POST" && path == "/data/bimbingan/mahasiswa":
		invokeLambdaHandler(w, r, "postBimbinganMahasiswaLambdaFunction")
	case method == "GET" && path == "/data/dosen":
		invokeLambdaHandler(w, r, "getDosenLambdaFunction")
	case method == "POST" && path == "/jadwalmengajar":
		invokeLambdaHandler(w, r, "getJadwalMengajarLambdaFunction")
	case method == "POST" && path == "/riwayatmengajar":
		invokeLambdaHandler(w, r, "getRiwayatPerkuliahanLambdaFunction")
	case method == "POST" && path == "/absensi":
		invokeLambdaHandler(w, r, "getAbsensiKelasLambdaFunction")
	case method == "POST" && path == "/nilai":
		invokeLambdaHandler(w, r, "getNilaiMahasiswaLambdaFunction")
	case method == "POST" && path == "/BAP":
		invokeLambdaHandler(w, r, "getBAPLambdaFunction")
	case method == "GET" && path == "/data/list/ta":
		invokeLambdaHandler(w, r, "getListTugasAkhirMahasiswaLambdaFunction")
	case method == "POST" && path == "/data/list/bimbingan":
		invokeLambdaHandler(w, r, "getListBimbinganMahasiswaLambdaFunction")
	case method == "POST" && path == "/approve/bimbingan":
		invokeLambdaHandler(w, r, "approveBimbinganLambdaFunction")
	default:
		http.NotFound(w, r)
	}
}

func invokeLambdaHandler(w http.ResponseWriter, r *http.Request, functionName string) {
	var requestBody map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	payload, err := json.Marshal(requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responsePayload, err := invokeLambda(functionName, payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responsePayload)
}
