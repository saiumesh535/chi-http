package utils

import (
	"fmt"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("TheSecretKey")

type jwtStructure struct {
    Data interface{} `json:"data"`
    jwt.StandardClaims
}

func init() {
	mySigningKey = []byte("TheSecretKey")
}

// CreateJwtToken token
func CreateJwtToken(input string) (string, error)  {
	claims := jwtStructure {
		input,
		jwt.StandardClaims {
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(8)).Unix(),
			Issuer:    "Yo!!",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}

// DecodeJwtToken decoding jwt token
func DecodeJwtToken(inputToken string) {
	token, err :=	jwt.Parse(inputToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
	} else {
		fmt.Println("err", err)
	}
}


