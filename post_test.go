package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostHandler_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(postHandler))
	defer ts.Close()

	bodyEcho := bytes.NewBufferString("Hello world!")
	respPost, err := http.Post(ts.URL, "application/json", bodyEcho)
	if err != nil {
		t.Fatal(err)
	}
	outPost, err := io.ReadAll(respPost.Body)
	if err != nil {
		t.Fatal(err)
	}
	stringOut := string(outPost)

	if stringOut != "Hello world!" {
		t.Fatalf(`Output should be "Hello world!",  but you have %s`, stringOut)
	}
}
func TestPostHandler_Wrong(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(postHandler))
	defer ts.Close()

	bodyEcho := bytes.NewBufferString("Hello world!")
	reqPost, err := http.NewRequest("GET", ts.URL, bodyEcho)
	if err != nil {
		t.Fatal(err)
	}
	client := &http.Client{}
	req, err := client.Do(reqPost)
	if err != nil {
		t.Fatal(err)
	}
	outPost, err := io.ReadAll(req.Body)
	if err != nil {
		t.Fatal(err)
	}

	var errPost errPost
	err = json.Unmarshal(outPost, &errPost)
	if err != nil {
		t.Fatal(err)
	}

	MsgErr := "Supports only POST method. Please use POST method."
	if errPost.ErrMsg != MsgErr {
		t.Fatalf("Error should be: %s\n but i recieved: %s\n", MsgErr, errPost.ErrMsg)
	}
}
