package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
)

func TestStatusHandler_Success(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/echo/{name}", echoHandler)

	tstServer := httptest.NewServer(r)
	defer tstServer.Close()

	resp, err := http.Post(tstServer.URL+"/echo/200", "content-type/text", nil)
	if err != nil {
		t.Fatal(err)
	}

	out, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	code, err := strconv.Atoi(string(out))
	if err != nil {
		t.Fatal(err)
	}
	if code != 200 {
		t.Fatalf(`Output should be "200",  but you have %d`, code)
	}
}
func TestStatusHandler_String(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/status/{codes}", statusHandler)

	tstServer := httptest.NewServer(r)
	defer tstServer.Close()

	resp, err := http.Post(tstServer.URL+"/status/dils", "content-type/text", nil)
	if err != nil {
		t.Fatal(err)
	}
	out, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var statusMsg statusResponse
	json.Unmarshal(out, &statusMsg)
	if statusMsg.ErrMsg != "Invalid status code. Code must be an integer. Example: 200, got: dils" {
		t.Fatalf(`Output should be an integer,  but you have %s`, statusMsg.ErrMsg)
	}
}

func TestStatusHandler_WronStatus(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/status/{codes}", statusHandler)

	tstServer := httptest.NewServer(r)
	defer tstServer.Close()

	resp, err := http.Post(tstServer.URL+"/status/199", "content-type/text", nil)
	if err != nil {
		t.Fatal(err)
	}
	out, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var statusMsg statusResponse
	json.Unmarshal(out, &statusMsg)
	if statusMsg.ErrMsg != "Invalid status code. Code is not a valid HTTP status code: 199" {
		t.Fatalf(`Output should be HTTP status code,  but you have %s`, statusMsg.ErrMsg)
	}
}
