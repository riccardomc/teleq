package client

import (
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestClient(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	t.Run("Test size", func(t *testing.T) {
		httpmock.RegisterResponder("GET", "http://localhost:9009/size",
			httpmock.NewStringResponder(200, `{"Operation":"size","Data":1000}`))

		client := TeleqClient{}
		response, err := client.Size("http://localhost:9009/")

		if err != nil {
			t.Error(err)
			return
		}
		if response != 1000 {
			t.Errorf("'%d' != %d", response, 1000)
		}
	})

	t.Run("Test empty", func(t *testing.T) {
		httpmock.RegisterResponder("GET", "http://localhost:9009/empty",
			httpmock.NewStringResponder(200, `{"Operation":"empty","Data":true}`))

		client := TeleqClient{}
		response, err := client.Empty("http://localhost:9009/")

		if err != nil {
			t.Error(err)
			return
		}
		if response != true {
			t.Errorf("'%t' != %t", response, true)
		}
	})

	t.Run("Test peek", func(t *testing.T) {
		httpmock.RegisterResponder("GET", "http://localhost:9009/peek",
			httpmock.NewStringResponder(200, `{"Operation":"peek","Data":"something"}`))

		client := TeleqClient{}
		response, err := client.Peek("http://localhost:9009/")

		if err != nil {
			t.Error(err)
			return
		}
		if response != "something" {
			t.Errorf("'%s' != %s", response, "something")
		}
	})

	t.Run("Test pop", func(t *testing.T) {
		httpmock.RegisterResponder("GET", "http://localhost:9009/pop",
			httpmock.NewStringResponder(200, `{"Operation":"pop","Data":"something"}`))

		client := TeleqClient{}
		response, err := client.Pop("http://localhost:9009/")

		if err != nil {
			t.Error(err)
			return
		}
		if response != "something" {
			t.Errorf("'%s' != %s", response, "something")
		}
	})

	t.Run("Test push", func(t *testing.T) {
		httpmock.RegisterResponder("POST", "http://localhost:9009/push",
			httpmock.NewStringResponder(200, `{"Operation":"push","Data":"something"}`))

		client := TeleqClient{}
		response, err := client.Push("http://localhost:9009/", "something")

		if err != nil {
			t.Error(err)
			return
		}
		if response != "something" {
			t.Errorf("'%s' != %s", response, "something")
		}
	})
}
