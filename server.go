package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

type msgErr struct {
	ErrMsg string
}

func main() {
	port := flag.String("port", "", "a string")
	flag.Parse()
	if *port == "" {
		fmt.Println("Please give a port number. Example --port 7077")
		return
	}

	http.HandleFunc("/echo", echoHandler)
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/ip", ipHandler)
	http.HandleFunc("/userAgent", userAgentHandler)
	http.HandleFunc("/headers", headersHandler)

	fmt.Printf("Starting server: %s\n", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", *port), nil)
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
