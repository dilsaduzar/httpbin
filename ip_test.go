package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIpHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(ipHandler))
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Host = "78.175.231.108"

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	out, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var ipJSON ip
	err = json.Unmarshal(out, &ipJSON)
	if err != nil {
		t.Fatal(err)
	}

	if ipJSON.Origin != "78.175.231.108" {
		t.Fatalf("ip result is wrong\n\nwant: %s\n got: %s\n", "78.175.231.108", ipJSON.Origin)
	}
}
