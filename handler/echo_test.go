package handler

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestEchoHandler(t *testing.T) {
	tstServer := httptest.NewServer(http.HandlerFunc(EchoHandler))
	defer tstServer.Close()

	bodyEcho := bytes.NewBufferString("Dilsad")

	respEcho, err := http.Post(tstServer.URL, "content-type/text", bodyEcho)
	if err != nil {
		t.Fatal(err)
	}
	outEcho, err := io.ReadAll(respEcho.Body)
	if err != nil {
		t.Fatal(err)
	}
	stringOut := string(outEcho)
	if stringOut != "Dilsad" {
		t.Fatalf(`Output should be "Dilsad",  but you have %s`, stringOut)
	}
}

func TestEchoHandler_Path(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/echo/{name}", EchoHandler)

	tstServer := httptest.NewServer(r)
	defer tstServer.Close()

	respEcho, err := http.Post(tstServer.URL+"/echo/Dilsad", "content-type/text", nil)
	if err != nil {
		t.Fatal(err)
	}
	outEcho, err := io.ReadAll(respEcho.Body)
	if err != nil {
		t.Fatal(err)
	}
	stringOut := string(outEcho)
	if stringOut != "Dilsad" {
		t.Fatalf(`Output should be "Dilsad",  but you have %s`, stringOut)
	}
}
