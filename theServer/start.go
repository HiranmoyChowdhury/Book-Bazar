package theServer

import (
	"github.com/go-chi/chi"
	"learnProject/First-Project-With-Go/api"
	"learnProject/First-Project-With-Go/utils"
	"net/http"
)

func Start() {
	r := chi.NewRouter()

	//api.Book calls all api related to book
	api.Book(r)

	err := http.ListenAndServe(utils.PortNo, r)
	if err != nil {
		return
	}
}
