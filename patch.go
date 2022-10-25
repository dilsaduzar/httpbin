package main

import (
	"io"
	"net/http"
)

type patchResponse struct {
	ErrMsg string
}

func patchHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPatch:
		out, err := io.ReadAll(r.Body)
		if err != nil {
			io.WriteString(w, errMsg("Error Code: PA01"))
			return
		}
		io.WriteString(w, string(out))
	default:
		w.WriteHeader(501)
		io.WriteString(w, errMsg("Supports only PATCH method. Please use PATCH method."))
		return
	}
}
