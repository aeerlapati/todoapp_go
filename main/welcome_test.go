package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func TestAddTask(t *testing.T) {
	t.Run("returns tweet", func(t *testing.T) {
		values := map[string]string{"taskname": "test1", "taskdescription": "test2"}
		jsonValue, _ := json.Marshal(values)
		requestBody := bytes.NewReader(jsonValue)

		request, _ := http.NewRequest(http.MethodPost, "/addTask", requestBody)
		response := httptest.NewRecorder()

		addTask(response, request)

		addtaskret := response.Body.String()
		fmt.Println(addtaskret)
		addtaskexp := "POST done, added task"

		if addtaskexp != addtaskexp {
			t.Errorf("got %q, want %q", addtaskret, addtaskexp)
		}
	})
}
