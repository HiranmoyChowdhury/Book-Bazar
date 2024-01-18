package utils

var PortNo string = ":3667"
var JWTSecretKey = []byte("SecretKey=*whatever*36")
var ExpTimeOfJWT = 1

func SetPortNo(port string) {
	PortNo = ":" + port
}
