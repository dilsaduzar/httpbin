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

	if len(out) > 0 {
		io.WriteString(w, string(out))
	} else if len(out) == 0 {
		vars := mux.Vars(r)
		key := vars["name"]

		if _, ok := vars["name"]; ok {
			io.WriteString(w, key)
		}
	}

}
