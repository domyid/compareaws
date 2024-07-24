package domyApi

import "github.com/aws/aws-lambda-go/events"

// Daftar origins yang diizinkan
var Origins = []string{
	"https://naskah.bukupedia.co.id",
	"https://chatgpl.do.my.id",
	"https://do.my.id",
	"https://in.my.id",
	"https://my.my.id",
	"https://whatsauth.github.io",
	"https://www.do.my.id",
}

// Fungsi untuk mengatur header CORS untuk Lambda
func SetAccessControlHeadersForLambda(request *events.APIGatewayProxyRequest) bool {
	origin := request.Headers["Origin"]

	if isAllowedOrigin(origin) {
		// Set CORS headers for the preflight request
		if request.HTTPMethod == "OPTIONS" {
			headers := map[string]string{
				"Access-Control-Allow-Credentials": "true",
				"Access-Control-Allow-Headers":     "Content-Type,Login",
				"Access-Control-Allow-Methods":     "POST,GET,DELETE,PUT",
				"Access-Control-Allow-Origin":      origin,
				"Access-Control-Max-Age":           "3600",
			}
			for key, value := range headers {
				request.Headers[key] = value
			}
			return true
		}
		// Set CORS headers for the main request.
		request.Headers["Access-Control-Allow-Credentials"] = "true"
		request.Headers["Access-Control-Allow-Origin"] = origin
		return false
	}

	return false
}

func isAllowedOrigin(string) bool {
	// Implementasi pengecekan apakah origin diizinkan
	// Contoh: return origin == "http://allowed-origin.com"
	return true // Ubah logika sesuai kebutuhan
}
