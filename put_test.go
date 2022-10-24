package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPutHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(putHandler))
	defer ts.Close()

	body := bytes.NewBufferString("Hello world!")
	req, err := http.NewRequest(http.MethodPut, ts.URL, body)
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

func TestPutHandler_Wrong(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(putHandler))
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

	var putResp errPut
	err = json.Unmarshal(out, &putResp)
	if err != nil {
		t.Fatal()
	}
	errMessage := "Supports only PUT method. Please use PUT method."
	if putResp.ErrMsg != errMessage {
		t.Fatalf("Error should be: %s\n but i recieved: %s\n", errMessage, putResp.ErrMsg)
	}
}
