package model

import (
	"time"
)

type User struct {
	Nipp      string    `json:"nipp" bson:"nipp"`
	Nama      string    `json:"nama" bson:"nama"`
	Jabatan   string    `json:"jabatan" bson:"jabatan"`
	Location  Location  `json:"location"`
	Password  string    `json:"password" bson:"password"`
	Role      string    `json:"role,omitempty" bson:"role,omitempty"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}

type Credential struct {
	Status  bool   `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
	Role    string `json:"role,omitempty" bson:"role,omitempty"`
}

type ResponseDataUser struct {
	Status  bool   `json:"status" bson:"status"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
	Data    []User `json:"data,omitempty" bson:"data,omitempty"`
}

type Response struct {
	Token string `json:"token,omitempty" bson:"token,omitempty"`
}

type ResponseEncode struct {
	Message string `json:"message,omitempty" bson:"message,omitempty"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
}

type Payload struct {
	User string    `json:"user"`
	Role string    `json:"role"`
	Exp  time.Time `json:"exp"`
	Iat  time.Time `json:"iat"`
	Nbf  time.Time `json:"nbf"`
}

type ResponseBack struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    []string `json:"data"`
}

type Location struct {
	LocationId   string `json:"locationId" bson:"locationId"`
	LocationName string `json:"locationName" bson:"locationName"`
}

type Area struct {
	AreaId   string `json:"areaId" bson:"areaId"`
	AreaName string `json:"areaName" bson:"areaName"`
}

type Cred struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ReqUsers struct {
	Nipp string `json:"nipp"`
}
