package main

import (
	"github.com/go-chi/chi"
	"learnProject/First-Project-With-Go/api"
	"learnProject/First-Project-With-Go/handler"
	"learnProject/First-Project-With-Go/utils"
	"net/http"
)

func main() {
	// r is the new Router for our server
	r := chi.NewRouter()

	allBookInfo := handler.GetBookList()

	//api.Book calls all api related to book
	api.Book(r, allBookInfo)

	err := http.ListenAndServe(utils.GetPort(), r)
	if err != nil {
		return
	}
}
