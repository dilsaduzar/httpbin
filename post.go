package main

import (
	"io"
	"net/http"
)

type errPost struct {
	ErrMsg string
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		out, err := io.ReadAll(r.Body)
		if err != nil {
			io.WriteString(w, errMsg("Error Code: P1"))
			return
		}
		io.WriteString(w, string(out))
	default:
		w.WriteHeader(501)
		io.WriteString(w, errMsg("Supports only POST method. Please use POST method."))
		return
	}
}
