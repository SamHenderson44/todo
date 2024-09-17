package handlers

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	store "github.com/SamHenderson44/todo/internal/storePackage"
)

func TestHandleGet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(HandleGet))
	testToDo := "testTo"
	store.GetStore().Add(testToDo)
	t.Run("has the correct response code", func(t *testing.T) {

		resp, err := http.Get(server.URL)

		if err != nil {
			t.Error(err)
		}

		assertStatusCode(t, resp.StatusCode, http.StatusOK)

	})

	t.Run("has the correct content type", func(t *testing.T) {

		resp, err := http.Get(server.URL)

		if err != nil {
			t.Error(err)
		}

		want := "text/html"

		contentType := resp.Header.Get("Content-Type")

		if !strings.HasPrefix(contentType, want) {
			t.Errorf("expected content type to to include %s but got %s ", want, contentType)
		}

	})

	tests := []struct {
		name     string
		expected string
	}{
		{"Correct ID", "<span>1</span>"},
		{"Correct Title", "<span>" + testToDo + "</span>"},
		{"Correct Completed status ", "<span>Not Done</span>"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.Get(server.URL)
			if err != nil {
				t.Error(err)
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Error(err)
			}

			html := string(body)
			if !strings.Contains(html, tt.expected) {
				t.Errorf("expected to find %s in the HTML but did not", tt.expected)
			}
		})
	}

}

func TestCreateToDoHandler(t *testing.T) {
	store := store.GetStore()

	t.Run("successfully adds a new to do", func(t *testing.T) {
		store.ResetStore()
		form := url.Values{}
		form.Add("toDo", "Test Todo")
		req := httptest.NewRequest(http.MethodPost, "/create", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(HandleCreateToDo)
		handler.ServeHTTP(rr, req)

		fmt.Println(store.GetToDos())

		assertStatusCode(t, rr.Code, http.StatusFound)
		assertToDoCount(t, len(store.ToDos), 1)
	})

	t.Run("responds with bad request if no to do value is added", func(t *testing.T) {
		store.ResetStore()
		form := url.Values{}
		form.Add("toDo", "")
		req := httptest.NewRequest(http.MethodPost, "/create", strings.NewReader(form.Encode()))

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(HandleCreateToDo)
		handler.ServeHTTP(rr, req)

		assertStatusCode(t, rr.Code, http.StatusBadRequest)
		assertToDoCount(t, len(store.ToDos), 0)
	})
}

func TestHandleUpdateStatus(t *testing.T) {
	t.Run("Updates status to true", func(t *testing.T) {})
	t.Run("Updates status to false", func(t *testing.T) {})
	t.Run("Returns bad request for invalid id format", func(t *testing.T) {})
	t.Run("Returns bad request when unable to parse request body", func(t *testing.T) {})
	t.Run("Returns not found for invalid id", func(t *testing.T) {})
}

func assertStatusCode(t *testing.T, statusCode int, expectedCode int) {
	if statusCode != expectedCode {
		t.Errorf("expected status code %d, got %d", expectedCode, statusCode)
	}
}

func assertToDoCount(t *testing.T, toDoCount int, expectedCount int) {
	t.Helper()
	if toDoCount != expectedCount {
		t.Errorf("expected %v to dos but got %v", expectedCount, toDoCount)
	}
}
