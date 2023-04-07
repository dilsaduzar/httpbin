package handler

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
type statusCode struct {
	StatusCode string
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
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

	resultJSON := statusCode{
		StatusCode: result,
	}
	out, err := json.Marshal(resultJSON)
	if err != nil {
		io.WriteString(w, errMsg("Error Code: S01"))
		return
	}
	fmt.Fprintln(w, string(out))
}
