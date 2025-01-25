package singlylinkedlist

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	checkSize := func(t *testing.T, list *List[string], expectedSize int) {
		t.Helper()
		if list.size() != expectedSize {
			t.Errorf("expected size %d, got %d", expectedSize, list.size())
		}
	}

	checkContents := func(t *testing.T, list *List[string], expected []string) {
		t.Helper()
		if list.size() != len(expected) {
			t.Errorf("expected size %d, got %d", len(expected), list.size())
			return
		}
		for i, v := range expected {
			if list.get(i) != v {
				t.Errorf("at index %d: expected %s, got %s", i, v, list.get(i))
			}
		}
	}

	expectPanic := func(t *testing.T, f func(), expectedMessage string) {
		t.Helper()
		defer func() {
			if r := recover(); r != nil {
				if r != expectedMessage {
					t.Errorf("unexpected panic message: got %v, want %v", r, expectedMessage)
				}
			} else {
				t.Errorf("expected panic but did not get one")
			}
		}()
		f()
	}

	t.Run("push_back and pop_back", func(t *testing.T) {
		list := &List[string]{}
		list.push_back("First")
		list.push_back("Second")
		list.push_back("Third")
		checkContents(t, list, []string{"First", "Second", "Third"})

		list.pop_back()
		checkContents(t, list, []string{"First", "Second"})

		list.pop_back()
		checkContents(t, list, []string{"First"})

		list.pop_back()
		checkContents(t, list, []string{})

		expectPanic(t, func() { list.pop_back() }, "attempt to pop_back from an empty list")
	})

	t.Run("push_front and pop_front", func(t *testing.T) {
		list := &List[string]{}
		list.push_front("Third")
		list.push_front("Second")
		list.push_front("First")
		checkContents(t, list, []string{"First", "Second", "Third"})

		list.pop_front()
		checkContents(t, list, []string{"Second", "Third"})

		list.pop_front()
		checkContents(t, list, []string{"Third"})

		list.pop_front()
		checkContents(t, list, []string{})

		expectPanic(t, func() { list.pop_front() }, "attempt to pop_front from an empty list")
	})

	t.Run("insert", func(t *testing.T) {
		list := &List[string]{}
		list.push_back("First")
		list.push_back("Third")
		list.insert("Second", 1)
		checkContents(t, list, []string{"First", "Second", "Third"})

		list.insert("Zero", 0)
		checkContents(t, list, []string{"Zero", "First", "Second", "Third"})

		list.insert("Fourth", 4)
		checkContents(t, list, []string{"Zero", "First", "Second", "Third", "Fourth"})

		expectPanic(t, func() { list.insert("Invalid", -1) }, "index -1 is out of range [0, 5]")
		expectPanic(t, func() { list.insert("Invalid", 6) }, "index 6 is out of range [0, 5]")
	})

	t.Run("removeAt", func(t *testing.T) {
		list := &List[string]{}
		list.push_back("First")
		list.push_back("Second")
		list.push_back("Third")
		list.push_back("Fourth")

		list.removeAt(1) // Remove "Second"
		checkContents(t, list, []string{"First", "Third", "Fourth"})

		list.removeAt(2) // Remove "Fourth"
		checkContents(t, list, []string{"First", "Third"})

		list.removeAt(0) // Remove "First"
		checkContents(t, list, []string{"Third"})

		list.removeAt(0) // Remove "Third"
		checkContents(t, list, []string{})

		expectPanic(t, func() { list.removeAt(-1) }, "index -1 is out of range [0, 0)")
		expectPanic(t, func() { list.removeAt(0) }, "index 0 is out of range [0, 0)")
	})

	t.Run("clear", func(t *testing.T) {
		list := &List[string]{}
		list.push_back("First")
		list.push_back("Second")
		list.push_back("Third")
		checkSize(t, list, 3)

		list.clear()
		checkSize(t, list, 0)
		checkContents(t, list, []string{})

		expectPanic(t, func() { list.get(0) }, "index 0 is out of range [0, 0)")
	})
}
