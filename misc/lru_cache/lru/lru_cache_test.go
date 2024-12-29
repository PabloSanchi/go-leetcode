package lru

import "testing"

func TestPutFirstElement(t *testing.T) {

	cache := NewLruCache[string, int](1)
	cache.Put("a", 1)

	if cache.head.key != "a" {
		t.Errorf("Expected head key to be 1, got %v", cache.head.key)
	}
	if cache.tail.key != "a" {
		t.Errorf("Expected tail key to be 1, got %v", cache.tail.key)
	}

	if cache.head.next != nil || cache.head.back != nil {
		t.Errorf("Expected head to be nil as is the only node, got %v", cache.head)
	}

}

func TestPutSecondElement(t *testing.T) {
	cache := NewLruCache[string, int](2)

	cache.Put("a", 1)
	cache.Put("b", 2)

	if cache.head.key != "b" {
		t.Errorf("Expected head key to be \"b\", got %v", cache.head.key)
	}

	if cache.tail.key != "a" {
		t.Errorf("Expected tail key to be \"a\", got %v", cache.tail.key)
	}
}

func TestGetElement(t *testing.T) {
	cache := NewLruCache[string, int](10)

	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	if value, _ := cache.Get("a"); value != 1 {
		t.Errorf("Expected value to be 1, got %v", value)
	}
}

func TestGetNotExistingElement(t *testing.T) {
	cache := NewLruCache[string, int](10)

	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	if _, ok := cache.Get("d"); ok {
		t.Errorf("Expected key \"d\" not to exist, but it does")
	}

}

func TestIsEmpty(t *testing.T) {
	cache := NewLruCache[string, int](10)

	if !cache.IsEmpty() {
		t.Errorf("Expected cache to be empty, but it is not")
	}

	cache.Put("a", 1)

	if cache.IsEmpty() {
		t.Errorf("Expected cache not to be empty, but it is")
	}
}

func TestSize(t *testing.T) {
	cache := NewLruCache[string, int](10)

	if cache.Size() != 10 {
		t.Errorf("Expected cache size to be 10, got %v", cache.Size())
	}
}
