package main

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	out, err := io.ReadAll(r.Body)
	if err != nil {
		io.WriteString(w, errMsg("Error Code: 01"))
		return
	}

	if string(out) != "" {
		io.WriteString(w, string(out))
	} else {
		vars := mux.Vars(r)
		if key, ok := vars["name"]; ok {
			io.WriteString(w, key)
		}
	}

}
