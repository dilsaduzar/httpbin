package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type msgErr struct {
	ErrMsg string
}

func main() {
	port := flag.String("port", "", "port number of the running server")
	flag.Parse()
	if *port == "" {
		fmt.Println("Please give a port number. Example --port 7077")
		return
	}

	r := mux.NewRouter()

	r.HandleFunc("/", rootHandler)

	r.HandleFunc("/echo/{name}", echoHandler)
	r.HandleFunc("/status/{codes}", statusHandler)

	// client requests
	r.HandleFunc("/ip", ipHandler)
	r.HandleFunc("/user-agent", userAgentHandler)
	r.HandleFunc("/headers", headersHandler)

	// methods
	r.HandleFunc("/get", getHandler)
	r.HandleFunc("/post", postHandler)
	r.HandleFunc("/put", putHandler)
	r.HandleFunc("/delete", deleteHandler)
	r.HandleFunc("/patch", patchHandler)

	fmt.Printf("Starting server: %s\n", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", *port), r)
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
