package api

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"learnProject/First-Project-With-Go/handler"
	"learnProject/First-Project-With-Go/model"
	"learnProject/First-Project-With-Go/utils"
	"net/http"
)

func Book(router *chi.Mux) {
	router.Use(MiddleWare)
	router.Post("/api/v1/books", handler.Add())
	router.Get("/api/v1/books", handler.GetAll())
	router.Get("/api/v1/books/{UUID}", handler.Get())
	router.Put("/api/v1/books/{UUID}", handler.Update())
	router.Delete("/api/v1/books/{UUID}", handler.Delete())

}
func MiddleWare(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if message, valid := utils.Authanticated(request); valid == false {
			writer.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(writer).Encode(model.Error{"401", "StatusUnauthorized", message})
			return
		}
		handler.ServeHTTP(writer, request)
	})
}
