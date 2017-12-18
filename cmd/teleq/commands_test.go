package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/urfave/cli"
)

func TestServer(t *testing.T) {
}

type clientTest struct {
	operation    string
	verb         string
	responseCode int
	responseBody string
	command      []string
	expected     string
}

func TestClient(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []clientTest{
		clientTest{
			"size",
			"GET",
			200,
			`{"Operation":"size","Data":1000}`,
			[]string{"", "size"},
			"1000\n",
		},
		clientTest{
			"push",
			"POST",
			200,
			`{"Operation":"push","Data":1000}`,
			[]string{"", "push", "-d", "1000"},
			"1000\n",
		},
		clientTest{
			"peek",
			"GET",
			200,
			`{"Operation":"peek","Data":1000}`,
			[]string{"", "peek"},
			"1000\n",
		},
		clientTest{
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

	t.Run("Test server", func(t *testing.T) {
		// Given
		outputBuffer := bytes.NewBuffer([]byte{})
		ServerAction = func(c *cli.Context) error {
			fmt.Fprint(c.App.Writer, "bla")
			return nil
		}
		app := New()

		// When
		app.Writer = outputBuffer
		err := app.Run([]string{"", "server"})

		// Then
		if err != nil {
			t.Error(err)
		}
		actualOutput := outputBuffer.String()
		expectedOutput := "bla"
		if actualOutput != expectedOutput {
			t.Errorf("'%s' != '%s'", actualOutput, expectedOutput)
		}
	})
}
