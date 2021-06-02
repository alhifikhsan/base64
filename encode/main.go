package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var kalimat string
	fmt.Print("Masukan kalimat ASCII:")
	fmt.Scanln(&kalimat)

	var encode = base64.StdEncoding.EncodeToString([]byte(kalimat))
	fmt.Println("Base64 :", encode)

	var decodeByte, _ = base64.StdEncoding.DecodeString(encode)
	var decode = string(decodeByte)
	fmt.Println("ASCII :", decode)
}
