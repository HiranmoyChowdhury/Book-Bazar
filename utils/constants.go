package utils

var PortNo string = ":3667"

func SetPortNo(port string) {
	PortNo = ":" + port
}
