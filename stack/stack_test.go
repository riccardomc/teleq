package stack

import "testing"

func TestNewStack(t *testing.T) {
	s := New()
	if s == nil {
		t.Error("Initialization failed")
		return
	}

	if s.Size() != 0 {
		t.Error("Created stack should be empty")
		return
	}
}

func TestPush(t *testing.T) {
	s := New()
	s.Push(1)
	if s.Size() != 1 {
		t.Errorf("Stack size should be 1 != %d", s.Size())
	}
	s.Push(2)
	if s.Size() != 2 {
		t.Errorf("Stack size should be 2 != %d", s.Size())
	}
	s.Push(3)
	if s.Size() != 3 {
		t.Errorf("Stack size should be 3 != %d", s.Size())
	}
}

func TestPeek(t *testing.T) {
	s := New()
	data := s.Peek()
	if data != nil {
		t.Errorf("data should be nil")
	}
	s.Push(1)
	data = s.Peek()
	if data != 1 {
		t.Errorf("data should be 1 != %d", data)
	}
	if s.Size() != 1 {
		t.Errorf("Size should be 1 != %d", s.Size())
	}
	s.Push(2)
	data = s.Peek()
	if data != 2 {
		t.Errorf("data should be 2 != %d", data)
	}
	if s.Size() != 2 {
		t.Errorf("Size should be 2 != %d", s.Size())
	}
}

func TestPop(t *testing.T) {
	s := New()

	t.Run("Pop on empty stack", func(t *testing.T) {
		data := s.Pop()
		if data != nil {
			t.Error("data should be nil")
		}
		if s.Size() != 0 {
			t.Errorf("size should be 0 != %d", s.Size())
		}

	})

	t.Run("Pop on stack with one element", func(t *testing.T) {
		s.Push(1)
		data := s.Pop()
		if data != 1 {
			t.Errorf("data should be 1 != %d", data)
		}
		if s.Size() != 0 {
			t.Errorf("size should be 0 != %d", s.Size())
		}
		if s.Peek() != nil {
			t.Errorf("top element should be nil")
		}
	})

	t.Run("Pop on stack with multiple elements", func(t *testing.T) {
		s.Push(1)
		s.Push(2)
		s.Push(3)
		data := s.Pop()
		if data != 3 {
			t.Errorf("data should be 3 != %d", data)
		}
		if s.Size() != 2 {
			t.Errorf("size should be 2 != %d", s.Size())
		}
		if s.Peek() != 2 {
			t.Errorf("top element should be 2 != %d", s.Peek())
		}
	})
}
