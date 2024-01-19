package auth

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"learnProject/First-Project-With-Go/model"
	"learnProject/First-Project-With-Go/utils"
	"net/http"
	"time"
)

var TokenForUser map[string]string = make(map[string]string)

func CreateJwtToken(userName string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": userName,
			"exp":      time.Now().Add(time.Second * 200).Unix(),
		})

	tokenString, err := token.SignedString([]byte(utils.JWTSecretKey))
	//fmt.Println("create token", userName, tokenString, err)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func VerifyToken(tokenString string, userName string) error {
	fmt.Println(tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.JWTSecretKey), nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	if value, present := TokenForUser[userName]; present == true && value == tokenString {
		return nil
	}

	return fmt.Errorf("Token Not Found In History")
}

func TokenValidation(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		userName := request.Header.Get("User")
		if token := request.Header.Get("Token"); VerifyToken(token, userName) != nil {
			json.NewEncoder(writer).Encode(model.Error{"401", "StatusUnauthorized", "Token is not valid"})
			return
		}
		handler.ServeHTTP(writer, request)

	})
}
func GetToken() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		userName, _, ok := request.BasicAuth()

		if !ok {
			writer.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(writer).Encode(model.Error{"401", "StatusUnauthorized", "user Name or Password is invalid"})
			return
		}
		if message, valid := Authanticated(request); valid == false {
			writer.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(writer).Encode(model.Error{"401", "StatusUnauthorized", message})
			return
		}

		token, _ := CreateJwtToken(userName)
		TokenForUser[userName] = token
		writer.Header().Add("Token", token)
		return

	}
}
