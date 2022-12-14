package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type headersResponse struct {
	Headers map[string]string `json:"headers,omitempty"`
}

func HeadersHandler(w http.ResponseWriter, r *http.Request) {
	headers := map[string]string{}
	for key, value := range r.Header {
		headers[key] = strings.Join(value, ",")
	}

	body := headersResponse{
		Headers: headers,
	}

	out, err := json.MarshalIndent(body, " ", "  ")
	if err != nil {
		io.WriteString(w, errMsg("Error Code: H01"))
		return
	}

	io.WriteString(w, string(out))
}
