package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPatchHandler_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(patchHandler))
	defer ts.Close()

	body := bytes.NewBufferString("Hello world!")
	req, err := http.NewRequest(http.MethodPatch, ts.URL, body)
	if err != nil {
		t.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	out, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	stringOut := string(out)

	if stringOut != "Hello world!" {
		t.Fatalf(`Output should be "Hello world!",  but you have %s`, stringOut)
	}
}
func TestPatchHandler_Wrong(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(patchHandler))
	defer ts.Close()

	body := bytes.NewBufferString("Hello world!")
	req, err := http.NewRequest(http.MethodGet, ts.URL, body)
	if err != nil {
		t.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	out, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	var msgPatch patchResponse
	err = json.Unmarshal(out, &msgPatch)
	if err != nil {
		t.Fatal(err)
	}
	errMsg := "Supports only PATCH method. Please use PATCH method."
	if msgPatch.ErrMsg != errMsg {
		t.Fatalf("Error should be: %s\n but i recieved: %s\n", errMsg, msgPatch.ErrMsg)
	}
}
