package handler

import (
	"encoding/json"
	"io"
	"net/http"
)

type errDelete struct {
	ErrMsg string
}

type msgErr struct {
	ErrMsg string
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "DELETE":
		out, err := io.ReadAll(r.Body)
		if err != nil {
			io.WriteString(w, errMsg("Error Code: D01"))
			return
		}
		io.WriteString(w, string(out))
	default:
		w.WriteHeader(501)
		io.WriteString(w, errMsg("Supports only DELETE method. Please use DELETE method."))
		return
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
