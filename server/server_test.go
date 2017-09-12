package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSize(t *testing.T) {
	targetServer := NewStackServer()
	request, err := http.NewRequest("GET", "/size", nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Size on Stack with no elements", func(t *testing.T) {
		expectedStatus := http.StatusOK
		expectedContent := "0"

		response, err := call(targetServer, request)
		if err != nil {
			t.Fatal(err)
		}

		actualStatus := response.Code
		actualContent := strings.TrimRight(response.Body.String(), "\n\t ")
		if actualStatus != expectedStatus {
			t.Errorf("Unexpected status: '%v'", actualStatus)
		}
		if actualContent != expectedContent {
			t.Errorf("Unexpected content: '%v'", actualContent)
		}
	})

	t.Run("Size on Stack with one element", func(t *testing.T) {
		targetServer.Stack.Push("one element")
		expectedStatus := http.StatusOK
		expectedContent := "1"

		response, err := call(targetServer, request)
		if err != nil {
			t.Fatal(err)
		}

		actualStatus := response.Code
		actualContent := strings.TrimRight(response.Body.String(), "\n\t ")
		if actualStatus != expectedStatus {
			t.Errorf("Unexpected status: '%v'", actualStatus)
		}
		if actualContent != expectedContent {
			t.Errorf("Unexpected content: '%v'", actualContent)
		}
	})

	t.Run("Size on Stack with multiple elements", func(t *testing.T) {
		targetServer.Stack.Push("another element")
		targetServer.Stack.Push("yet another element")
		expectedStatus := http.StatusOK
		expectedContent := "3"

		response, err := call(targetServer, request)
		if err != nil {
			t.Fatal(err)
		}

		actualStatus := response.Code
		actualContent := strings.TrimRight(response.Body.String(), "\n\t ")
		if actualStatus != expectedStatus {
			t.Errorf("Unexpected status: '%v'", actualStatus)
		}
		if actualContent != expectedContent {
			t.Errorf("Unexpected content: '%v'", actualContent)
		}
	})
}

func TestPeek(t *testing.T) {
	targetServer := NewStackServer()
	request, err := http.NewRequest("GET", "/peek", nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Peek on Stack with no elements", func(t *testing.T) {
		expectedStatus := http.StatusOK
		expectedContent := "<nil>"

		response, err := call(targetServer, request)
		if err != nil {
			t.Fatal(err)
		}

		actualStatus := response.Code
		if actualStatus != expectedStatus {
			t.Errorf("Unexpected status: '%v'", actualStatus)
		}
		actualContent := strings.TrimRight(response.Body.String(), "\n\t ")
		if actualContent != expectedContent {
			t.Errorf("Unexpected content: '%v'", actualContent)
		}
	})

	t.Run("Peek on Stack with one element", func(t *testing.T) {
		targetServer.Stack.Push("one element")
		expectedStatus := http.StatusOK
		expectedContent := "one element"

		response, err := call(targetServer, request)
		if err != nil {
			t.Fatal(err)
		}

		actualStatus := response.Code
		if actualStatus != expectedStatus {
			t.Errorf("Unexpected status: '%v'", actualStatus)
		}
		actualContent := strings.TrimRight(response.Body.String(), "\n\t ")
		if actualContent != expectedContent {
			t.Errorf("Unexpected content: '%v'", actualContent)
		}
	})

	t.Run("Peek on Stack with multiple elements", func(t *testing.T) {
		targetServer.Stack.Push("another element")
		expectedStatus := http.StatusOK
		expectedContent := "another element"

		response, err := call(targetServer, request)
		if err != nil {
			t.Fatal(err)
		}

		actualStatus := response.Code
		if actualStatus != expectedStatus {
			t.Errorf("Unexpected status: '%v'", actualStatus)
		}
		actualContent := strings.TrimRight(response.Body.String(), "\n\t ")
		if actualContent != expectedContent {
			t.Errorf("Unexpected content: '%v'", actualContent)
		}
	})
}

func TestPush(t *testing.T) {
	targetServer := NewStackServer()

	t.Run("Push on Stack with no elements", func(t *testing.T) {
		request, _ := http.NewRequest("POST", "/push/one element", nil)
		expectedStatus := http.StatusOK
		expectedContent := "one element"

		response, err := call(targetServer, request)
		if err != nil {
			t.Fatal(err)
		}

		actualStatus := response.Code
		if actualStatus != expectedStatus {
			t.Errorf("Unexpected status: '%v'", actualStatus)
		}
		actualContent := strings.TrimRight(response.Body.String(), "\n\t ")
		if actualContent != expectedContent {
			t.Errorf("Unexpected content: '%v'", actualContent)
		}
		actualStackContent := targetServer.Stack.Peek()
		if actualStackContent != expectedContent {
			t.Errorf("Unexpected stack content: '%v'", actualStackContent)
		}
		actualStackSize := targetServer.Stack.Size()
		if actualStackSize != 1 {
			t.Errorf("Unexpected stack size: '%v'", actualStackSize)
		}
	})

	t.Run("Push on Stack with one element", func(t *testing.T) {
		request, _ := http.NewRequest("POST", "/push/another element", nil)
		expectedStatus := http.StatusOK
		expectedContent := "another element"

		response, err := call(targetServer, request)
		if err != nil {
			t.Fatal(err)
		}

		actualStatus := response.Code
		if actualStatus != expectedStatus {
			t.Errorf("Unexpected status: '%v'", actualStatus)
		}
		actualContent := strings.TrimRight(response.Body.String(), "\n\t ")
		if actualContent != expectedContent {
			t.Errorf("Unexpected content: '%v'", actualContent)
		}
		actualStackContent := targetServer.Stack.Peek()
		if actualStackContent != expectedContent {
			t.Errorf("Unexpected stack content: '%v'", actualStackContent)
		}
		actualStackSize := targetServer.Stack.Size()
		if actualStackSize != 2 {
			t.Errorf("Unexpected stack size: '%v'", actualStackSize)
		}
	})

	t.Run("Push on Stack with multiple elements", func(t *testing.T) {
		request, _ := http.NewRequest("POST", "/push/yet another element", nil)
		expectedStatus := http.StatusOK
		expectedContent := "yet another element"

		response, err := call(targetServer, request)
		if err != nil {
			t.Fatal(err)
		}

		actualStatus := response.Code
		if actualStatus != expectedStatus {
			t.Errorf("Unexpected status: '%v'", actualStatus)
		}
		actualContent := strings.TrimRight(response.Body.String(), "\n\t ")
		if actualContent != expectedContent {
			t.Errorf("Unexpected content: '%v'", actualContent)
		}
		actualStackContent := targetServer.Stack.Peek()
		if actualStackContent != expectedContent {
			t.Errorf("Unexpected stack content: '%v'", actualStackContent)
		}
		actualStackSize := targetServer.Stack.Size()
		if actualStackSize != 3 {
			t.Errorf("Unexpected stack size: '%v'", actualStackSize)
		}
	})
}

func TestPop(t *testing.T) {
	targetServer := NewStackServer()
	request, err := http.NewRequest("GET", "/pop", nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Pop on Stack with no elements", func(t *testing.T) {
		expectedStatus := http.StatusOK
		expectedContent := "<nil>"

		response, err := call(targetServer, request)
		if err != nil {
			t.Fatal(err)
		}

		actualStatus := response.Code
		if actualStatus != expectedStatus {
			t.Errorf("Unexpected status: '%v'", actualStatus)
		}
		actualContent := strings.TrimRight(response.Body.String(), "\n\t ")
		if actualContent != expectedContent {
			t.Errorf("Unexpected content: '%v'", actualContent)
		}
	})

	t.Run("Pop on Stack with one element", func(t *testing.T) {
		targetServer.Stack.Push("one element")
		expectedStatus := http.StatusOK
		expectedContent := "one element"

		response, err := call(targetServer, request)
		if err != nil {
			t.Fatal(err)
		}

		actualStatus := response.Code
		if actualStatus != expectedStatus {
			t.Errorf("Unexpected status: '%v'", actualStatus)
		}
		actualContent := strings.TrimRight(response.Body.String(), "\n\t ")
		if actualContent != expectedContent {
			t.Errorf("Unexpected content: '%v'", actualContent)
		}
		actualStackContent := targetServer.Stack.Peek()
		if actualStackContent != nil {
			t.Errorf("Unexpected stack content: '%v'", actualStackContent)
		}
		actualStackSize := targetServer.Stack.Size()
		if actualStackSize != 0 {
			t.Errorf("Unexpected stack size: '%v'", actualStackSize)
		}
	})

	t.Run("Pop on Stack with multiple elements", func(t *testing.T) {
		targetServer.Stack.Push("one element")
		targetServer.Stack.Push("another element")
		expectedStatus := http.StatusOK
		expectedContent := "another element"

		response, err := call(targetServer, request)
		if err != nil {
			t.Fatal(err)
		}

		actualStatus := response.Code
		if actualStatus != expectedStatus {
			t.Errorf("Unexpected status: '%v'", actualStatus)
		}
		actualContent := strings.TrimRight(response.Body.String(), "\n\t ")
		if actualContent != expectedContent {
			t.Errorf("Unexpected content: '%v'", actualContent)
		}
		actualStackContent := targetServer.Stack.Peek()
		if actualStackContent != "one element" {
			t.Errorf("Unexpected stack content: '%v'", actualStackContent)
		}
		actualStackSize := targetServer.Stack.Size()
		if actualStackSize != 1 {
			t.Errorf("Unexpected stack size: '%v'", actualStackSize)
		}
	})
}

func call(targetServer *StackServer, request *http.Request) (*httptest.ResponseRecorder, error) {
	method := request.Method
	path := request.URL.String()
	targetHandle, params, _ := targetServer.Router.Lookup(method, path)
	if targetHandle == nil {
		return nil, fmt.Errorf("Cannot find route '%s' '%s'", method, path)
	}

	handler := http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) { targetHandle(w, r, params) })
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)

	return recorder, nil
}
