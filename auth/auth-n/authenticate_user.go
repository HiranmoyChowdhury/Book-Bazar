package auth_n

import (
	"learnProject/First-Project-With-Go/model"
	"net/http"
)

func Authanticated(r *http.Request) (string, bool) {
	username, password, ok := r.BasicAuth()
	if ok == false {
		return "header value is Invalid", false
	}
	if model.UserNameMatchPass(username, password) == true {
		return "valid user", true
	}
	return "user name or password is incorrect", false

}
