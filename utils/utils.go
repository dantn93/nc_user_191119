package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang191119/nc_user/config"
	"github.com/golang191119/nc_user/model"
	"time"
)

func MD5(text string) string {
	data := []byte(text)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func GenerateToken(userID int, phone, email string) string {
	signKey := []byte(config.Config.JWTSecret.JWTKey)
	fmt.Println("secret", config.Config.JWTSecret.JWTKey, "end secret")
	claims := model.UserClaims{
		UserID: userID,
		Phone:  phone,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(48 * time.Hour).Unix(), //unix timestamp
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(signKey)

	return ss
}
