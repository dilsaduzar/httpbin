package handler

import (
	"io"
	"net/http"
)

type errPut struct {
	ErrMsg string
}

func PutHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		out, err := io.ReadAll(r.Body)
		if err != nil {
			io.WriteString(w, errMsg("Error Code: PU01"))
			return
		}
		io.WriteString(w, string(out))
	default:
		w.WriteHeader(501)
		io.WriteString(w, errMsg("Supports only PUT method. Please use PUT method."))
		return
	}
}
