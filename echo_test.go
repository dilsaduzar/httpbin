package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEchoHansler(t *testing.T) {
	tstServer := httptest.NewServer(http.HandlerFunc(echoHandler))

	bodyEcho := bytes.NewBufferString("Dilsad")

	respEcho, err := http.Post(tstServer.URL, "content-type/ text", bodyEcho)
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
