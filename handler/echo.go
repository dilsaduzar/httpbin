package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type msgErr struct {
	ErrMsg string
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
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

func errMsg(msg string) string {
	Msg := msgErr{msg}
	outErr, err := json.Marshal(&Msg)
	if err != nil {
		return `{"Error code": "-1"}`
	}
	return string(outErr)
}
