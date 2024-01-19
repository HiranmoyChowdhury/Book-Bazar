package server

import (
	"github.com/go-chi/chi"
	"learnProject/First-Project-With-Go/api"
	"net/http"
)

func Start(portNo string) *chi.Mux {
	r := chi.NewRouter()

	api.Book(r)

	http.ListenAndServe(":"+portNo, r)
	return r
}
