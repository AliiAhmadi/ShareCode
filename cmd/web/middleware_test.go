package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSecureHeaders(t *testing.T) {
	responseRecorder := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	next := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("OK"))
	})

	secureHeaders(next).ServeHTTP(responseRecorder, request)
	result := responseRecorder.Result()

	frameOptions := result.Header.Get("X-Frame-Options")
	if frameOptions != "deny" {
		t.Errorf("want: %q; got: %q", "deny", frameOptions)
	}

	xssProtection := result.Header.Get("X-XSS-Protection")
	if xssProtection != "1; mode=block" {
		t.Errorf("want: %q; got: %q", "1; mode=block", xssProtection)
	}

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
