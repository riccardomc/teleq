package main

import (
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestServer(t *testing.T) {
}

func TestClient(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	t.Run("Test Size", func(t *testing.T) {
		// Given
		httpmock.RegisterResponder("GET", "http://localhost:9009/size",
			httpmock.NewStringResponder(200, `{"Operation":"size","Data":1000}`))
		// When
		app := New()
		err := app.Run([]string{"", "-a", "http://localhost:9009/", "size"})

		if err != nil {
			t.Error(err)
		}
	})
}
