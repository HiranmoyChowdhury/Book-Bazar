package auth_z

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"learnProject/First-Project-With-Go/utils"
	"time"
)

var TokenForUser map[string]string = make(map[string]string)

func CreateJwtToken(userName string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": userName,
			"exp":      time.Now().Add(time.Second * 200).Unix(),
		})

	tokenString, err := token.SignedString(utils.JWTSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func VerifyToken(tokenString string, userName string) error {
	fmt.Println(tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return utils.JWTSecretKey, nil
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
