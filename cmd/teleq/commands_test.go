package main

import (
	"bytes"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/julienschmidt/httprouter"
)

type MockServer struct{}

func (s *MockServer) Size() httprouter.Handle { return nil }
func (s *MockServer) Peek() httprouter.Handle { return nil }
func (s *MockServer) Push() httprouter.Handle { return nil }
func (s *MockServer) Pop() httprouter.Handle  { return nil }
func (s *MockServer) Serve(port int)          {}

type clientTest struct {
	operation    string
	verb         string
	responseCode int
	responseBody string
	command      []string
	expected     string
}

type serverTest struct {
	description string
	command     []string
	expected    string
}

func TestClient(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []clientTest{
		{
			"size",
			"GET",
			200,
			`{"Operation":"size","Data":1000}`,
			[]string{"", "size"},
			"1000\n",
		},
		{
			"push",
			"POST",
			200,
			`{"Operation":"push","Data":1000}`,
			[]string{"", "push", "-d", "1000"},
			"1000\n",
		},
		{
			"peek",
			"GET",
			200,
			`{"Operation":"peek","Data":1000}`,
			[]string{"", "peek"},
			"1000\n",
		},
		{
			"pop",
			"GET",
			200,
			`{"Operation":"pop","Data":1000}`,
			[]string{"", "pop"},
			"1000\n",
		},
	}

	for _, test := range tests {
		t.Run("Test "+test.operation, func(t *testing.T) {
			// Given
			httpmock.RegisterResponder(test.verb, "http://localhost:9009/"+test.operation,
				httpmock.NewStringResponder(test.responseCode, test.responseBody))
			outputBuffer := bytes.NewBuffer([]byte{})

			// When
			app := New()
			app.Writer = outputBuffer
			err := app.Run(test.command)

			// Then
			if err != nil {
				t.Error(err)
			}
			actualOutput := outputBuffer.String()
			if actualOutput != test.expected {
				t.Errorf("'%s' != '%s'", actualOutput, test.expected)
			}
		})

	}

}

func TestServer(t *testing.T) {
	tests := []serverTest{
		{
			"default port",
			[]string{"", "server"},
			"Serving on 9009\n",
		},
		{
			"custom port",
			[]string{"", "server", "-p", "9999"},
			"Serving on 9999\n",
		},
	}

	Server = &MockServer{}
	outputBuffer := bytes.NewBuffer([]byte{})

	for _, test := range tests {
		t.Run("Test server "+test.description, func(t *testing.T) {
			// Given
			app := New()

			// When
			app.Writer = outputBuffer
			err := app.Run(test.command)

			// Then
			if err != nil {
				t.Error(err)
			}
			actual := outputBuffer.String()
			if actual != test.expected {
				t.Errorf("'%s' != '%s'", actual, test.expected)
			}
			outputBuffer.Reset()
		})
	}
}
