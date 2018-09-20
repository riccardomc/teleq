package database

import (
	"bytes"
	"testing"
)

func TestFrame(t *testing.T) {
	t.Run("Serialize", func(t *testing.T) {
		f := Frame{"this is some data"}
		expectedSerialized := []byte{28, 255, 129, 3, 1, 1, 5, 70, 114, 97, 109, 101, 1, 255, 130, 0, 1, 1, 1, 4, 68, 97, 116, 97, 1, 16, 0, 0, 0, 32, 255, 130, 1, 6, 115, 116, 114, 105, 110, 103, 12, 19, 0, 17, 116, 104, 105, 115, 32, 105, 115, 32, 115, 111, 109, 101, 32, 100, 97, 116, 97, 0}
		actualSerialized, err := f.Serialize()
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(expectedSerialized, actualSerialized) {
			t.Error(expectedSerialized, "!=", actualSerialized)
		}
	})

	t.Run("Deserialize", func(t *testing.T) {
		data := []byte{28, 255, 129, 3, 1, 1, 5, 70, 114, 97, 109, 101, 1, 255, 130, 0, 1, 1, 1, 4, 68, 97, 116, 97, 1, 16, 0, 0, 0, 32, 255, 130, 1, 6, 115, 116, 114, 105, 110, 103, 12, 19, 0, 17, 116, 104, 105, 115, 32, 105, 115, 32, 115, 111, 109, 101, 32, 100, 97, 116, 97, 0}
		expectedDeserialized := Frame{"this is some data"}
		actualDeserialized := Frame{}
		_, err := actualDeserialized.Deserialize(data)
		if err != nil {
			t.Fatal(err)
		}
		if expectedDeserialized != actualDeserialized {
			t.Error(expectedDeserialized, "!=", actualDeserialized)
		}
	})
}
