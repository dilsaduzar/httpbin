package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEchoHandler(t *testing.T) {
	tstServer := httptest.NewServer(http.HandlerFunc(echoHandler))
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
	t.Skip()

	tstServer := httptest.NewServer(http.HandlerFunc(echoHandler))
	defer tstServer.Close()

	respEcho, err := http.Post(tstServer.URL+"/Dilsad", "content-type/text", nil)
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
