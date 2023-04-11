package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"

	"github.com/dilsaduzar/httpbin/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	port := flag.String("port", "", "port number of the running server")
	dburl := flag.String("sqluser", "", "user information of the database")
	flag.Parse()

	if *port == "" || *dburl == "" {
		fmt.Println(`Please give a port number and user information. Example --port 7077 --sqluser "example:123(127.0.0.1:3306)/example"`)

		return
	}

	db, err := sql.Open("mysql", *dburl)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	handler.DB = db

	r := mux.NewRouter()

	r.HandleFunc("/", handler.RootHandler)

	r.HandleFunc("/echo/{name}", handler.EchoHandler)
	r.HandleFunc("/status/{codes}", handler.StatusHandler)

	r.HandleFunc("/cities/{city}", handler.CitiesHandler)

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
	err = http.ListenAndServe(fmt.Sprintf(":%s", *port), r)
	if err != nil {
		fmt.Printf("Couldn`t start server. Error %s\n", err)
	}
}
