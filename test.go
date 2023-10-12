package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestImgcatGetRequest(t *testing.T) {
	// Create a mock HTTP server to simulate the request and response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate a successful response with a 200 status code
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Make a test request to the imgcatGetRequest function
	response := imgcatGetRequest(server.URL, server.Client())

	// Check if the response status code is 200 (OK)
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, but got %d", response.StatusCode)
	}
}
