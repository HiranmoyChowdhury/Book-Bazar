package model

var userInfo map[string]string = map[string]string{
	"H":  makePassFifty("Hydrogen"),
	"He": makePassFifty("Helium"),
}

func makePassFifty(pass string) string {
	ret := pass
	for len(ret) < 50 {
		ret += "1"
	}
	return ret
}
func UserNameMatchPass(name string, pass string) bool {
	if result, present := userInfo[name]; result == makePassFifty(pass) && present == true {
		return true
	}
	return false
}
