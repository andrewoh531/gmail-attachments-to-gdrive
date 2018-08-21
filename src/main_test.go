package main

import (
	//"fmt"
	//"net/http"
	//"net/http/httptest"
	"testing"

		)

func TestHandler(t *testing.T) {

	//// TODO This is an integration test (Ideally should stub/mock out calls to AWS)
	//t.Run("Should retrieve SSM", func(t *testing.T) {
	//
	//	response := retrieveFromParameterStore("andrew1")
	//	fmt.Print(response)
	//	_, err := handler(events.CloudWatchEvent{})
	//	if err == nil {
	//		t.Fatal("Error failed to trigger with an invalid request")
	//	}
	//})


	//t.Run("Unable to get IP", func(t *testing.T) {
	//	DefaultHTTPGetAddress = "http://127.0.0.1:12345"
	//
	//	_, err := handler(events.CloudWatchEvent{})
	//	if err == nil {
	//		t.Fatal("Error failed to trigger with an invalid request")
	//	}
	//})
	//
	//t.Run("Non 200 Response", func(t *testing.T) {
	//	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//		w.WriteHeader(500)
	//	}))
	//	defer ts.Close()
	//
	//	DefaultHTTPGetAddress = ts.URL
	//
	//	_, err := handler(events.CloudWatchEvent{})
	//	if err != nil && err.Error() != ErrNon200Response.Error() {
	//		t.Fatalf("Error failed to trigger with an invalid HTTP response: %v", err)
	//	}
	//})
	//
	//t.Run("Unable decode IP", func(t *testing.T) {
	//	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//		w.WriteHeader(500)
	//	}))
	//	defer ts.Close()
	//
	//	DefaultHTTPGetAddress = ts.URL
	//
	//	_, err := handler(events.CloudWatchEvent{})
	//	if err == nil {
	//		t.Fatal("Error failed to trigger with an invalid HTTP response")
	//	}
	//})
	//
	//t.Run("Successful Request", func(t *testing.T) {
	//	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//		w.WriteHeader(200)
	//		fmt.Fprintf(w, "127.0.0.1")
	//	}))
	//	defer ts.Close()
	//
	//	DefaultHTTPGetAddress = ts.URL
	//
	//	_, err := handler(events.CloudWatchEvent{})
	//	if err != nil {
	//		t.Fatal("Everything should be ok")
	//	}
	//})
}
