package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ip struct {
	Origin string `json:"origin,omitempty"`
}

func ipHandler(w http.ResponseWriter, r *http.Request) {
	ipInf := ip{
		Origin: r.Host,
	}

	ipOut, err := json.MarshalIndent(ipInf, "", "  ")
	if err != nil {
		io.WriteString(w, errMsg("Error Code: IP01"))
		return
	}

	fmt.Fprintf(w, "%s", string(ipOut))
}
