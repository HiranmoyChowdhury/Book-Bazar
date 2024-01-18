package api

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"learnProject/First-Project-With-Go/auth/auth-z"
	"learnProject/First-Project-With-Go/handler"
	"learnProject/First-Project-With-Go/model"
	"net/http"
)

func Book(router *chi.Mux) {

	router.Group(func(router chi.Router) {
		router.Use(TokenValidation)
		router.Post("/api/v1/books", handler.Add())
		router.Get("/api/v1/books", handler.GetAll())
		router.Get("/api/v1/books/{UUID}", handler.Get())
		router.Put("/api/v1/books/{UUID}", handler.Update())
		router.Delete("/api/v1/books/{UUID}", handler.Delete())
	})
	router.Group(func(router chi.Router) {
		router.Get("/api/v1/get-token", handler.GetToken())
	})

}
func TokenValidation(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		userName := request.Header.Get("User")
		if token := request.Header.Get("Token"); auth_z.VerifyToken(token, userName) != nil {
			json.NewEncoder(writer).Encode(model.Error{"401", "StatusUnauthorized", "Token is not valid"})
			return
		}
		handler.ServeHTTP(writer, request)

	})
}
