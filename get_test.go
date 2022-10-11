package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHandler_Success(t *testing.T) {
	tstGet := httptest.NewServer(http.HandlerFunc(getHandler))
	defer tstGet.Close()

	bodyEcho := bytes.NewBufferString("Hello world!")

	respGet, err := http.NewRequest("GET", "http://127.0.0.1:7070/get", bodyEcho)
	if err != nil {
		t.Fatal(err)
	}

	outGet, err := io.ReadAll(respGet.Body)
	if err != nil {
		t.Fatal(err)
	}
	stringOut := string(outGet)
	if stringOut != "Hello world!" {
		t.Fatalf(`Output should be "Hello world!",  but you have %s`, stringOut)
	}
}

func TestGetHandler_Wrong(t *testing.T) {
	tstGet := httptest.NewServer(http.HandlerFunc(getHandler))
	defer tstGet.Close()

	bodyEcho := bytes.NewBufferString("Hello world!")

	respPost, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:7070/get", bodyEcho)
	if err != nil {
		t.Fatal(err)
	}
	outPost, err := io.ReadAll(respPost.Body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(outPost))
	//Hata donmuyor. 38. satirda gitdigim degeri basiyor

	var errPost errResult
	err = json.Unmarshal(outPost, &errPost)
	if err != nil {
		t.Fatal(err)
	}
	if errPost.ErrMsg != "Supported only GET method. Please use GET method." {
		t.Fatalf("jghlaurg")
	}
}
