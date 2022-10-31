package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/dilsaduzar/httpbin/handler"
	"github.com/gorilla/mux"
)

func main() {
	port := flag.String("port", "", "port number of the running server")
	flag.Parse()
	if *port == "" {
		fmt.Println("Please give a port number. Example --port 7077")
		return
	}

	r := mux.NewRouter()

	r.HandleFunc("/", handler.RootHandler)

	r.HandleFunc("/echo/{name}", handler.EchoHandler)
	r.HandleFunc("/status/{codes}", handler.StatusHandler)

	// client requests
	r.HandleFunc("/ip", handler.IpHandler)
	r.HandleFunc("/user-agent", handler.UserAgentHandler)
	r.HandleFunc("/headers", handler.HeadersHandler)

	// methods
	r.HandleFunc("/get", handler.GetHandler)
	r.HandleFunc("/post", handler.PostHandler)
	r.HandleFunc("/put", handler.PutHandler)
	r.HandleFunc("/delete", handler.DeleteHandler)
	r.HandleFunc("/patch", handler.PatchHandler)

	fmt.Printf("Starting server: %s\n", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", *port), r)
	if err != nil {
		fmt.Printf("Couldn`t start server. Error %s\n", err)
	}
}
