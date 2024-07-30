package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"

	"github.com/domyid/helper"
	"github.com/domyid/model"
)

func Login(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var resp model.Credential
	var datauser model.User
	err := json.Unmarshal([]byte(request.Body), &datauser)
	if err != nil {
		resp.Message = "error parsing application/json: " + err.Error()
		return responseJSON(http.StatusBadRequest, resp), nil
	}

	mongoEnv := os.Getenv("MongoEnv")
	dbName := os.Getenv("DBName")
	colName := os.Getenv("ColName")
	privateKey := os.Getenv("PrivateKey")

	fmt.Println("MongoEnv:", mongoEnv)
	fmt.Println("DBName:", dbName)
	fmt.Println("ColName:", colName)
	fmt.Println("PrivateKey:", privateKey)

	mconn, err := helper.SetConnection(mongoEnv, dbName)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		resp.Message = "error connecting to database"
		return responseJSON(http.StatusInternalServerError, resp), nil
	}

	fmt.Println("Validating user password")
	if helper.PasswordValidator(mconn, colName, datauser) {
		fmt.Println("Password validation successful")
		datarole := helper.GetOneUser(mconn, colName, model.User{Nipp: datauser.Nipp})
		fmt.Println("User role retrieved:", datarole.Role)
		tokenstring, err := helper.EncodeWithRole(datarole.Role, datauser.Nipp, privateKey)
		if err != nil {
			resp.Message = "Gagal Encode Token : " + err.Error()
		} else {
			resp.Status = true
			resp.Token = tokenstring
			resp.Message = "Selamat Datang di Portsafe+"
			resp.Role = datarole.Role
		}
	} else {
		fmt.Println("Password validation failed")
		resp.Message = "Password Salah"
	}
	return responseJSON(http.StatusOK, resp), nil
}

func responseJSON(status int, body interface{}) events.APIGatewayProxyResponse {
	response, _ := json.Marshal(body)
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(response),
	}
}
