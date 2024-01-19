package auth

import (
	"encoding/json"
	"learnProject/First-Project-With-Go/model"
	"net/http"
)

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
