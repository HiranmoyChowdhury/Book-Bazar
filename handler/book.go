package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/oklog/ulid/v2"
	"learnProject/First-Project-With-Go/model"
	"net/http"
)

var theBookList []model.Book = make([]model.Book, 0, 1)
var locationOfBook map[string]int = make(map[string]int)

func GetBookList() *[]model.Book {
	return &theBookList
}
func GetABook() *model.Book {
	return &model.Book{}
}

func Add() http.HandlerFunc {
	memory := &theBookList
	return func(w http.ResponseWriter, r *http.Request) {
		//if message, valid := utils.Authanticated(r); valid == false {
		//	w.WriteHeader(http.StatusUnauthorized)
		//	json.NewEncoder(w).Encode(model.Error{"401", "StatusUnauthorized", message})
		//	return
		//}
		bookRecieved := GetABook()
		err := json.NewDecoder(r.Body).Decode(bookRecieved)
		if err != nil {
			json.NewEncoder(w).Encode(model.Error{"400", "Bad Request", fmt.Sprint(err)})
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if IsValid(bookRecieved) == false {
			json.NewEncoder(w).Encode(model.Error{"400", "Bad Request", "file format is not granted"})
			w.WriteHeader(400)
			return
		}
		bookRecieved.UUID = base64.StdEncoding.EncodeToString(ulid.Make().Bytes())
		locationOfBook[bookRecieved.UUID] = len(*memory)
		*memory = append(*memory, *bookRecieved)

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(bookRecieved)
	}
}

func Delete() http.HandlerFunc {
	memory := &theBookList
	return func(w http.ResponseWriter, r *http.Request) {

		//if message, valid := utils.Authanticated(r); valid == false {
		//	w.WriteHeader(http.StatusUnauthorized)
		//	json.NewEncoder(w).Encode(model.Error{"401", "StatusUnauthorized", message})
		//	return
		//}

		param := chi.URLParam(r, "UUID")
		if _, present := locationOfBook[param]; present == false {
			json.NewEncoder(w).Encode(model.Error{"404", "Not Found", "the file is missing!"})
			w.WriteHeader(404)
			return
		}
		idx := locationOfBook[param]
		bookList := *memory
		err := json.NewEncoder(w).Encode(bookList[idx])
		if err != nil {
			json.NewEncoder(w).Encode(model.Error{"500", "server issue", fmt.Sprint(err)})
			w.WriteHeader(500)
		}
		delete(locationOfBook, bookList[idx].UUID)
		if idx > 0 {
			theBookList = append(bookList[:idx], bookList[idx+1:]...)
		} else {
			theBookList = bookList[idx+1:]
		}
		w.WriteHeader(200)
	}
}

func Get() http.HandlerFunc {
	memory := &theBookList
	return func(w http.ResponseWriter, r *http.Request) {

		//if message, valid := utils.Authanticated(r); valid == false {
		//	w.WriteHeader(http.StatusUnauthorized)
		//	json.NewEncoder(w).Encode(model.Error{"401", "StatusUnauthorized", message})
		//	return
		//}
		param := chi.URLParam(r, "UUID")
		if _, present := locationOfBook[param]; present == false {
			json.NewEncoder(w).Encode(model.Error{"404", "not found", "file is missing"})
			w.WriteHeader(404)
			return
		}
		idx := locationOfBook[param]
		bookList := *memory
		err := json.NewEncoder(w).Encode(bookList[idx])
		if err != nil {
			json.NewEncoder(w).Encode(model.Error{"500", "server issue", fmt.Sprint(err)})
			w.WriteHeader(500)
		}
		w.WriteHeader(200)
	}
}

func GetAll() http.HandlerFunc {
	memory := &theBookList
	return func(w http.ResponseWriter, r *http.Request) {

		//if message, valid := utils.Authanticated(r); valid == false {
		//	w.WriteHeader(http.StatusUnauthorized)
		//	json.NewEncoder(w).Encode(model.Error{"401", "StatusUnauthorized", message})
		//	return
		//}
		err := json.NewEncoder(w).Encode(*memory)
		if err != nil {
			json.NewEncoder(w).Encode(model.Error{"500", "server issue", fmt.Sprint(err)})
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(200)

	}
}

func Update() http.HandlerFunc {
	memory := &theBookList
	return func(w http.ResponseWriter, r *http.Request) {

		//if message, valid := utils.Authanticated(r); valid == false {
		//	w.WriteHeader(http.StatusUnauthorized)
		//	json.NewEncoder(w).Encode(model.Error{"401", "StatusUnauthorized", message})
		//	return
		//}
		param := chi.URLParam(r, "UUID")

		if _, present := locationOfBook[param]; present == false {
			json.NewEncoder(w).Encode(model.Error{"404", "not found", "file is missing"})
			w.WriteHeader(404)
			return
		}
		bookList := *memory
		idx := locationOfBook[param]
		bookInfo := GetABook()
		err := json.NewDecoder(r.Body).Decode(bookInfo)
		if IsValid(bookInfo) == false {
			json.NewEncoder(w).Encode(model.Error{"404", "Bad Request", "file format is not granted"})
			w.WriteHeader(404)
			return
		}
		bookInfo.UUID = param
		if err != nil {
			json.NewEncoder(w).Encode(model.Error{"400", "Bad Request", fmt.Sprint(err)})
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		bookList[idx] = *bookInfo
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(bookInfo)

	}
}
