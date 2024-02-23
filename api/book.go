package api

import (
	"github.com/go-chi/chi"
	"learnProject/First-Project-With-Go/auth"
	"learnProject/First-Project-With-Go/handler"
)

func Book(router *chi.Mux) {

	router.Group(func(router chi.Router) {
		router.Use(auth.TokenValidation)
		router.Post("/api/v1/book", handler.Add())
		router.Get("/api/v1/books", handler.GetAll())
		router.Get("/api/v1/book/{UUID}", handler.Get())
		router.Put("/api/v1/book/{UUID}", handler.Update())
		router.Delete("/api/v1/book/{UUID}", handler.Delete())
	})
	router.Group(func(router chi.Router) {
		router.Post("/api/v1/get-token", auth.GetToken())
	})

}
