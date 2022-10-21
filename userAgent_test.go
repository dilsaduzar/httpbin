package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserAgent(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(userAgentHandler))
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("User-Agent", "Golang")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	out, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var userJSON userResponse
	err = json.Unmarshal(out, &userJSON)
	if err != nil {
		t.Fatal(err)
	}

	if userJSON.UserAgent != "Golang" {
		t.Fatalf("user agent result is wrong\n\nwant: %s\n got: %s\n", "Golang", userJSON.UserAgent)
	}
}
