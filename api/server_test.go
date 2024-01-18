package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
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
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		//panic("your server is not ready for it's use")
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestStart(t *testing.T) {
	// Create a New Server Struct
	s := chi.NewRouter()

	//http.ListenAndServe(utils.PortNo, s)

	f := testModelS{"sldkfj", []string{"lkjsdf"}, "sldfj", "kljsdf"}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(f)
	Book(s)

	// log in
	req, _ := http.NewRequest("GET", "/api/v1/get-token", &buf)
	fmt.Println("this part is done")
	req.SetBasicAuth("H", "hi")

	response := executeRequest(req, s)
	token := response.Header().Get("Token")

	fmt.Println("hi", token)
	fmt.Println("Post body: \n", response.Body)

	//fmt.Println("Post header: \n", response.Header())
	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", "/api/v1/books", &buf)
	fmt.Println("this part is done")
	req.Header.Set("Token", token)
	req.Header.Set("User", "H")

	response = executeRequest(req, s)

	fmt.Println("Post body: \n", response.Body)

	//fmt.Println("Post header: \n", response.Header())
	checkResponseCode(t, http.StatusOK, response.Code)

}
