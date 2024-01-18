package the_server

import (
	"github.com/go-chi/chi"
	"learnProject/First-Project-With-Go/api"
	"learnProject/First-Project-With-Go/utils"
	"net/http"
)

func Start() *chi.Mux {
	r := chi.NewRouter()

	api.Book(r)

	http.ListenAndServe(utils.PortNo, r)
	return r
}
