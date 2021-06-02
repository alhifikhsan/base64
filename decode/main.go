package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var kalimat string
	fmt.Print("Masukan kalimat Base64:")
	fmt.Scanln(&kalimat)

	var decodeByte, _ = base64.StdEncoding.DecodeString(kalimat)
	var decode = string(decodeByte)
	fmt.Println("Base64 :", decode)
}
