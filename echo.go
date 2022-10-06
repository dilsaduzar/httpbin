package main

import (
	"io"
	"net/http"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	out, err := io.ReadAll(r.Body)
	if err != nil {
		io.WriteString(w, errMsg("Error Code: 01"))
		return
	}
	io.WriteString(w, string(out))

}
