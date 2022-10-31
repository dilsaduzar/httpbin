package handler

import (
	"io"
	"net/http"
)

type errResult struct {
	ErrMsg string
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		out, err := io.ReadAll(r.Body)
		if err != nil {
			io.WriteString(w, errMsg("Error Code: A"))
			return
		}
		io.WriteString(w, string(out))
	default:
		w.WriteHeader(501)
		io.WriteString(w, errMsg("Supported only GET method. Please use GET method."))
		return
	}

}
