package main

import (
	"bytes"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestServer(t *testing.T) {
}

func TestClient(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	t.Run("Test size", func(t *testing.T) {
		// Given
		expectedOutput := "1000\n"
		httpmock.RegisterResponder("GET", "http://localhost:9009/size",
			httpmock.NewStringResponder(200, `{"Operation":"size","Data":1000}`))
		outputBuffer := bytes.NewBuffer([]byte{})

		// When
		app := New()
		app.Writer = outputBuffer
		err := app.Run([]string{"", "size"})

		// Then
		if err != nil {
			t.Error(err)
		}
		actualOutput := outputBuffer.String()
		if actualOutput != expectedOutput {
			t.Errorf("'%s' != '%s'", actualOutput, expectedOutput)
		}
	})
}
