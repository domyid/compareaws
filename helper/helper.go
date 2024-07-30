package helper

import (
	"encoding/json"
	"fmt"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	"github.com/domyid/model"
)

func AWSReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}

func SetConnection(MongoString, dbname string) (*mongo.Database, error) {
	fmt.Println("Connecting to MongoDB with URI:", MongoString)
	MongoInfo := atdb.DBInfo{
		DBString: MongoString,
		DBName:   dbname,
	}
	conn := atdb.MongoConnect(MongoInfo)
	if conn == nil {
		return nil, fmt.Errorf("error connecting to MongoDB: connection is nil")
	}
	fmt.Println("Connected to MongoDB")
	return conn, nil
}

func GetOneUser(MongoConn *mongo.Database, colname string, userdata model.User) model.User {
	filter := bson.M{"nipp": userdata.Nipp}
	data := atdb.GetOneDoc[model.User](MongoConn, colname, filter)
	fmt.Println("User data retrieved:", data)
	return data
}

func PasswordValidator(MongoConn *mongo.Database, colname string, userdata model.User) bool {
	fmt.Println("Starting password validation")
	filter := bson.M{"nipp": userdata.Nipp}
	data := atdb.GetOneDoc[model.User](MongoConn, colname, filter)
	fmt.Println("User data for validation:", data)
	hashChecker := CheckPasswordHash(userdata.Password, data.Password)
	fmt.Println("Password validation result:", hashChecker)
	return hashChecker
}

func EncodeWithRole(role, nipp, privatekey string) (string, error) {
	token := paseto.NewToken()
	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(2 * time.Hour))
	token.SetString("user", nipp)
	token.SetString("role", role)
	key, err := paseto.NewV4AsymmetricSecretKeyFromHex(privatekey)
	return token.V4Sign(key, nil), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
