package database

import (
	"os/exec"
	"testing"

	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/sqlite" //noqa
)

func TestStack(t *testing.T) {

	stack := New()
	database := "/tmp/" + uuid.New().String()
	defer func() {
		exec.Command("rm", "-f", database).Run()
	}()

	t.Run("Init", func(t *testing.T) {
		err := stack.Init(database)
		if err != nil {
			t.Error(err)
		}
		for _, table := range []string{"stacks", "persistent_frames"} {
			if !stack.db.client.HasTable(table) {
				t.Error("Table", table, "not initialized")
			}
		}
	})

	t.Run("Size", func(t *testing.T) {
		stack.Push("this")
		size := stack.Size()
		if !(size > 0) {
			t.Error(size, "!> 0")
		}
	})

	t.Run("Push", func(t *testing.T) {
		expectedSize := stack.Size() + 1
		stack.Push("something")
		actualSize := stack.Size()
		if expectedSize != actualSize {
			t.Error(actualSize, "!=", expectedSize)
		}
	})

	t.Run("Peek", func(t *testing.T) {
		expectedData := "a value"
		stack.Push(expectedData)
		actualData := stack.Peek()
		if actualData != expectedData {
			t.Error(actualData, "!=", expectedData)
		}
	})

	t.Run("Pop", func(t *testing.T) {
		expectedSize := stack.Size()
		expectedData := "this value"
		stack.Push(expectedData)

		actualData := stack.Pop()
		actualSize := stack.Size()

		if actualData != expectedData {
			t.Error(actualData, "!=", expectedData)
		}

		if actualSize != expectedSize {
			t.Error(actualSize, "!=", expectedSize)
		}
	})
}
