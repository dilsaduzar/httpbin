package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRoot(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(rootHandler))
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL, nil)
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
	result := string(out)

	tt := time.Now()
	resultRoot := fmt.Sprintln("Hello World!\n", tt.Format("02-01-2006 15:04:05"))

	if result != resultRoot {
		t.Fatalf("root result is wrong\n\nwant: %s\n got: %s\n", result, resultRoot)
	}
}
