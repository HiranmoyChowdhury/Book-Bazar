package api

import (
	"github.com/go-chi/chi"
	"learnProject/First-Project-With-Go/handler"
)

func Book(router *chi.Mux) {
	router.Post("/api/v1/books", handler.Add())
	router.Get("/api/v1/books", handler.GetAll())
	router.Get("/api/v1/books/{UUID}", handler.Get())
	router.Put("/api/v1/books/{UUID}", handler.Update())
	router.Delete("/api/v1/books/{UUID}", handler.Delete())
}
