package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHandler_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(GetHandler))
	defer ts.Close()

	bodyEcho := bytes.NewBufferString("Hello world!")

	respGet, err := http.NewRequest("GET", ts.URL, bodyEcho)
	if err != nil {
		t.Fatal(err)
	}

	respGet.Header.Set("Content-Type", "text/html")
	client := &http.Client{}
	req, err := client.Do(respGet)
	if err != nil {
		t.Fatal(err)
	}

	outGet, err := io.ReadAll(req.Body)
	if err != nil {
		t.Fatal(err)
	}
	stringOut := string(outGet)
	if stringOut != "Hello world!" {
		t.Fatalf(`Output should be "Hello world!",  but you have %s`, stringOut)
	}
}

func TestGetHandler_Wrong(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(GetHandler))
	defer ts.Close()

	bodyEcho := bytes.NewBufferString("Hello world!")

	respPost, err := http.NewRequest(http.MethodPost, ts.URL, bodyEcho)
	if err != nil {
		t.Fatal(err)
	}

	respPost.Header.Set("Content-Type", "text/html")
	client := &http.Client{}
	req, err := client.Do(respPost)
	if err != nil {
		t.Fatal(err)
	}

	outPost, err := io.ReadAll(req.Body)
	if err != nil {
		t.Fatal(err)
	}

	var errPost errResult
	err = json.Unmarshal(outPost, &errPost)
	if err != nil {
		t.Fatal(err)
	}
	MsgErr := "Supported only GET method. Please use GET method."
	if errPost.ErrMsg != MsgErr {
		t.Fatalf("Error should be: %s\n but i recieved: %s\n", MsgErr, errPost.ErrMsg)
	}
}
