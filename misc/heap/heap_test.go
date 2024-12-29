package heap

import "testing"

func TestHeapInsert(t *testing.T) {
	h := NewHeap[int]()
	h.Add(1)

	if h.Size() != 1 {
		t.Errorf("Expected size 1, got %d", h.Size())
	}

}

func TestOrderedInsert(t *testing.T) {
	h := NewHeap[int]()
	h.Add(3)
	h.Add(2)
	h.Add(1)

	if h.Size() != 3 {
		t.Errorf("Expected size 3, got %d", h.Size())
	}

	value, ok := h.Peek()
	if !ok {
		t.Errorf("Expected a value")
	}

	if value != 1 {
		t.Errorf("Expected 1, got %d", value)
	}
}

func TestHeapRemove(t *testing.T) {
	h := NewHeap[int]()
	h.Add(3)
	h.Add(2)
	h.Add(1)

	if h.Size() != 3 {
		t.Errorf("Expected size 3, got %d", h.Size())
	}

	_, ok := h.Delete()
	if !ok {
		t.Errorf("Expected success")
	}

	value, ok := h.Peek()
	if !ok {
		t.Errorf("Expected a value")
	}

	if value != 2 {
		t.Errorf("Expected 2, got %d", value)
	}
}
