package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCanary(t *testing.T) {
	t.Run("returns tweet", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/canary", nil)
		response := httptest.NewRecorder()

		Canary(response, request)

		got := response.Body.String()
		want := "tweet"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
