package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	responseRecorder := httptest.NewRecorder()

	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	ping(responseRecorder, request)
	result := responseRecorder.Result()

	if result.StatusCode != http.StatusOK {
		t.Errorf("want: %d; got: %d", http.StatusOK, result.StatusCode)
	}

	defer result.Body.Close()
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "OK" {
		t.Errorf("want: %s; got: %s", "OK", string(body))
	}
}
