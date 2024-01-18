package utils

var PortNo string = ":3667"
var JWTSecretKey = []byte("SecretKey=*whatever*36")

func SetPortNo(port string) {
	PortNo = ":" + port
}
