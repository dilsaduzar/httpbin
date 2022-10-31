package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteHandler_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(deleteHandler))
	defer ts.Close()

	body := bytes.NewBufferString("Hello world!")
	req, err := http.NewRequest("DELETE", ts.URL, body)
	if err != nil {
		t.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	outGet, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	stringOut := string(outGet)
	if stringOut != "Hello world!" {
		t.Fatalf(`Output should be "Hello world!",  but you have %s`, stringOut)
	}
}
func TestDeleteHandler_Wrong(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(deleteHandler))
	defer ts.Close()

	body := bytes.NewBufferString("Hello world!")
	req, err := http.NewRequest("GET", ts.URL, body)
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

	var deleteErr errDelete
	err = json.Unmarshal(out, &deleteErr)
	if err != nil {
		t.Fatal(err)
	}
	MsgErr := "Supports only DELETE method. Please use DELETE method."
	if deleteErr.ErrMsg != MsgErr {
		t.Fatalf("Error should be: %s\n but i recieved: %s\n", MsgErr, deleteErr.ErrMsg)
	}
}
