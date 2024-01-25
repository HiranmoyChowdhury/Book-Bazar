package auth

var userInfo map[string]string = map[string]string{
	"H":  "hi",
	"He": "he",
}

func UserNameMatchPass(name string, pass string) bool {
	//fmt.Println(name, pass, userInfo[name])
	if result, present := userInfo[name]; result == pass && present == true {
		return true
	}
	return false
}
func SetUserInfo(name string, pass string) {
	userInfo[name] = pass
}
