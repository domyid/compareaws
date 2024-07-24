package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	route "github.com/domyid/domyapi/route"
)

func main() {
	// Create a new HTTP server mux
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/login", route.URL)
	mux.HandleFunc("/refresh-token", route.URL)
	mux.HandleFunc("/data/mahasiswa", route.URL)
	mux.HandleFunc("/data/bimbingan/mahasiswa", route.URL)
	mux.HandleFunc("/data/dosen", route.URL)
	mux.HandleFunc("/jadwalmengajar", route.URL)
	mux.HandleFunc("/riwayatmengajar", route.URL)
	mux.HandleFunc("/absensi", route.URL)
	mux.HandleFunc("/nilai", route.URL)
	mux.HandleFunc("/BAP", route.URL)
	mux.HandleFunc("/data/list/ta", route.URL)
	mux.HandleFunc("/data/list/bimbingan", route.URL)
	mux.HandleFunc("/approve/bimbingan", route.URL)

	// Create an HTTP adapter
	adapter := httpadapter.New(mux)

	// Start the Lambda function
	lambda.Start(adapter.Proxy)
}
