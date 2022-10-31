package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type statusResponse struct {
	ErrMsg string
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["codes"]

	status, err := strconv.Atoi(code)
	if err != nil {
		w.WriteHeader(400)
		io.WriteString(w, errMsg(fmt.Sprintf("Invalid status code. Code must be an integer. Example: 200, got: %s", code)))
		return
	}

	result := http.StatusText(status)
	if result == "" {
		w.WriteHeader(400)
		io.WriteString(w, errMsg(fmt.Sprintf("Invalid status code. Code is not a valid HTTP status code: %s", code)))
		return
	}

	w.WriteHeader(status)
}
func errMsg(msg string) string {
	Msg := msgErr{msg}
	outErr, err := json.Marshal(&Msg)
	if err != nil {
		return `{"Error code": "-1"}`
	}
	return string(outErr)
}
