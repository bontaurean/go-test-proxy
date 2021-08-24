package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bontaurean/go-test-proxy/models"
)

func TestChain(t *testing.T) {
	e := httptest.NewServer(setupServer(true))
	defer e.Close()

	req, _ := json.Marshal(&models.ProxyRequest{
		Method:  "GET",
		URL:     "https://postman-echo.com/get",
		Headers: models.PlainHeaders{"Authentication": "Basic bG9naW46cGFzc3dvcmQ="},
	})

	res, err := http.Post(fmt.Sprintf("%s/v1/requests", e.URL), "application/json", bytes.NewBuffer(req))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", res.StatusCode)
	}

	ct, ok := res.Header["Content-Type"]

	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	if ct[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected 'application/json; charset=utf-8', got %s", ct[0])
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		t.Fatal("Could not read POST response body")
	}

	// ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

	pResp := models.ProxyResponse{}
	if err := json.Unmarshal(body, &pResp); err != nil {
		t.Fatal("Could not parse POST response body as JSON")
	}

	res, _ = http.Get(fmt.Sprintf("%s/v1/requests/%s", e.URL, pResp.ID))

	if res.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", res.StatusCode)
	}

	ct, ok = res.Header["Content-Type"]

	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	if ct[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected 'application/json; charset=utf-8', got %s", ct[0])
	}

	body, readErr = ioutil.ReadAll(res.Body)
	if readErr != nil {
		t.Fatal("Could not read GET response body")
	}

	hEntry := models.HistoryEntry{}
	if err := json.Unmarshal(body, &hEntry); err != nil {
		t.Fatal("Could not parse GET response body as JSON")
	}

	t.Logf("%#v", hEntry)
}
