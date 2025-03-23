package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	t.Run("Returns Hello World", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/hello", nil)
		response := httptest.NewRecorder()
		HelloServer(response, request)
		got := response.Body.String()
		want := "Hello World"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

// type StubUserStore struct {
// 	input string
// }
//
// func TestCli(t *testing.T) {
// 	userstore := &StubUserStore{}
// 	cli := &CLI{userstore}
// 	cli.RecordInput()
//
// 	if (userstore.input != "horses") {
// 		t.Fatal("expected user input to be horses")
// 	}
// }
