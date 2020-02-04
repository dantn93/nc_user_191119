package model

import (
	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name" bson:"first_name,omitempty"`
	LastName  string `json:"last_name" bson:"last_name,omitempty"`
	Email     string `json:"email" bson:"email,omitempty"`
	Phone     string `json:"phone" bson:"phone,omitempty"`
	Password  string `json:"password" bson:"password,omitempty"`
}

type UserClaims struct {
	UserID int    `json:"uid"`
	Phone  string `json:"p"`
	Email  string `json:"e"`
	jwt.StandardClaims
}

type Error struct {
	Code int
	Msg  string
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginResp struct {
	*User
	Token string `json:"token"`
}

type Counters struct {
	ID            string `json:"id,omitempty"`
	SequenceValue int    `json:"sequence_value,omitempty"`
}
