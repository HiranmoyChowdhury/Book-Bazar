package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"learnProject/First-Project-With-Go/handler"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testModelS struct {
	Name        string   `json:"name"`
	AuthorList  []string `json:"authorList"`
	PublishDate string   `json:"publishDate"`
	ISBN        string   `json:"isbn"`
}

func executeRequest(req *http.Request, s *chi.Mux) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.ServeHTTP(rr, req)

	return rr
}

// checkResponseCode is a simple utility to check the response code
// of the response
func checkResponseCode(t *testing.T, expected int, response *httptest.ResponseRecorder, testBody testModelS) {

	bookRecieved := handler.GetABook()
	err := json.NewDecoder(response.Body).Decode(bookRecieved)
	if err != nil {
		t.Errorf("body format is not valid")
	} else if response.Code != expected {
		t.Errorf("Expected response code %d. Got %d\n", expected, response.Code)
	} else if bookRecieved.Name != testBody.Name {
		t.Errorf("error in name")
	} else if bookRecieved.PublishDate != testBody.PublishDate {
		t.Errorf("error in Publish date")
	} else if bookRecieved.ISBN != testBody.ISBN {
		t.Errorf("error in ISBN")
	} else if len(bookRecieved.AuthorList) != len(testBody.AuthorList) {
		t.Errorf("error in author list")
	} else {
		for i := 0; i < len(bookRecieved.AuthorList); i++ {
			if bookRecieved.AuthorList[i] != testBody.AuthorList[i] {
				t.Errorf("error in author list")
				return
			}
		}
	}

}

func TestAdd(t *testing.T) {
	// Create a New Server Struct
	s := chi.NewRouter()
	Book(s)

	//http.ListenAndServe(utils.PortNo, s)

	testBody := testModelS{"sldkfj", []string{"lkjsdf"}, "sldfj", "kljsdf"}
	var testBuf bytes.Buffer
	json.NewEncoder(&testBuf).Encode(testBody)

	// log in
	req, _ := http.NewRequest("GET", "/api/v1/get-token", nil)
	fmt.Println("log in part is done")
	req.SetBasicAuth("H", "hi")

	response := executeRequest(req, s)
	token := response.Header().Get("Token")
	fmt.Println(response.Code)

	req, _ = http.NewRequest("POST", "/api/v1/book", &testBuf)

	fmt.Println("post req part is done")
	req.Header.Set("Token", token)
	req.Header.Set("User", "H")

	response = executeRequest(req, s)

	checkResponseCode(t, http.StatusOK, response, testBody)

}
