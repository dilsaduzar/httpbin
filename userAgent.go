package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type userResponse struct {
	UserAgent string `json:"user-agent"`
}

func userAgentHandler(w http.ResponseWriter, r *http.Request) {
	uAgent := userResponse{
		UserAgent: r.UserAgent(),
	}

	uOut, err := json.MarshalIndent(uAgent, "", "  ")
	if err != nil {
		io.WriteString(w, errMsg("Error Code: UA01"))
		return
	}
	io.WriteString(w, string(uOut))
}
