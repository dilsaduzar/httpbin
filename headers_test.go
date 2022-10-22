package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHeaderHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(headersHandler))
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header = map[string][]string{
		"Accept": {"A", "B"},
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
	fmt.Println(string(out))

	var headersTest headersResponse
	err = json.Unmarshal(out, &headersTest)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(headersTest.Headers)

	if headersTest.Headers["Accept"] != "A,B" {
		t.Fatalf("headers result is wrong\n\nwant: %s\n got: %s\n", "A,B", headersTest.Headers["Accept"])
	}

}
