package api

import (
	"github.com/go-chi/chi"
	"learnProject/First-Project-With-Go/handler"
	"learnProject/First-Project-With-Go/model"
)

func Book(router *chi.Mux, allBookInfo *[]model.Book) {
	//create a Book
	router.Post("/api/v1/books", handler.Add(allBookInfo))
	//get all books
	router.Get("/api/v1/books", handler.GetAll(allBookInfo))
	// get a book by book id
	router.Get("/api/v1/books/{UUID}", handler.Get(allBookInfo))
	//update book by bookId
	router.Put("/api/v1/books/{UUID}", handler.Update(allBookInfo))
	//Delete a book by bookId
	router.Delete("/api/v1/books/{UUID}", handler.Delete(allBookInfo))
}
