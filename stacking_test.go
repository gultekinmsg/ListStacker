package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)
func Test501(t *testing.T){
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(quoteHandler)

	body, err := json.Marshal(ExpectedQuotes{
		{AUTHOR: "MSG", QUOTE: "hi"},
		{AUTHOR: "EA", QUOTE: "go"},
		{AUTHOR: "MFG", QUOTE: "postman"},
		{AUTHOR: "MFG", QUOTE: "trello"},
		{AUTHOR: "EA", QUOTE: "vue"},
		{AUTHOR: "EA", QUOTE: "flutter"},
		{AUTHOR: "EA", QUOTE: "html"},
	})
	if err != nil {
		t.Fatal(err)
	}
	read := bytes.NewReader(body)
	request, err := http.NewRequest(http.MethodGet, "/quote", read)
	if err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(recorder, request)
	status := recorder.Code
	if status != http.StatusNotImplemented {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
func Test404 (t *testing.T){
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(quoteHandler)

	body, err := json.Marshal(ExpectedQuotes{
		{AUTHOR: "MSG", QUOTE: "hi"},
		{AUTHOR: "EA", QUOTE: "go"},
		{AUTHOR: "MFG", QUOTE: "postman"},
		{AUTHOR: "MFG", QUOTE: "trello"},
		{AUTHOR: "EA", QUOTE: "vue"},
		{AUTHOR: "EA", QUOTE: "flutter"},
		{AUTHOR: "EA", QUOTE: "html"},
	})
	if err != nil {
		t.Fatal(err)
	}
	read := bytes.NewReader(body)
	request, err := http.NewRequest(http.MethodPost, "/quotexx", read)
	if err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(recorder, request)
	status := recorder.Code
	if status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
func TestHandler(t *testing.T) {
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(quoteHandler)

	body, err := json.Marshal(ExpectedQuotes{
		{AUTHOR: "MSG", QUOTE: "hi"},
		{AUTHOR: "EA", QUOTE: "go"},
		{AUTHOR: "MFG", QUOTE: "postman"},
		{AUTHOR: "MFG", QUOTE: "trello"},
		{AUTHOR: "EA", QUOTE: "vue"},
		{AUTHOR: "EA", QUOTE: "flutter"},
		{AUTHOR: "EA", QUOTE: "html"},
	})
	if err != nil {
		t.Fatal(err)
	}
	read := bytes.NewReader(body)
	request, err := http.NewRequest(http.MethodPost, "/quote", read)
	if err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(recorder, request)
	status := recorder.Code
	if status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := []GivenQuotes{
		{AUTHOR: "MSG", QUOTES: getSlices("hi")},
		{AUTHOR: "EA", QUOTES: getSlices("go", "vue", "flutter", "html")},
		{AUTHOR: "MFG", QUOTES: getSlices("postman", "trello")},
	}
	var result []GivenQuotes
	err = json.NewDecoder(recorder.Body).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("we got %s result should be %s", result, expected)
	}
}
func getSlices(quotes ...string) []string {
	var slice []string
	for i := 0; i < len(quotes); i++ {
		slice = append(slice, quotes[i])
	}
	return slice
}
