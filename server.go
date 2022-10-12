package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type msgErr struct {
	ErrMsg string
}

func main() {
	http.HandleFunc("/echo", echoHandler)
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/post", postHandler)
	fmt.Println("Starting server: 7070..")
	err := http.ListenAndServe(":7070", nil)
	if err != nil {
		fmt.Printf("Couldn`t start server. Error %s\n", err)
	}
}
func errMsg(msg string) string {
	Msg := msgErr{msg}
	outErr, err := json.Marshal(&Msg)
	if err != nil {
		return `{"Error code": "-1"}`
	}
	return string(outErr)
}
