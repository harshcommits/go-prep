package syntaxtest

import "testing"

func TestLRUCache(t *testing.T) {
	cache := NewLRUCache(2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	if got := cache.Get(1); got != 1 {
		t.Fatalf("expected value 1 for key 1, got %d", got)
	}

	cache.Put(3, 3)

	if got := cache.Get(2); got != -1 {
		t.Fatalf("expected miss for key 2, got %d", got)
	}

	cache.Put(4, 4)

	if got := cache.Get(1); got != -1 {
		t.Fatalf("expected miss for key 1, got %d", got)
	}

	if got := cache.Get(3); got != 3 {
		t.Fatalf("expected value 3 for key 3, got %d", got)
	}

	if got := cache.Get(4); got != 4 {
		t.Fatalf("expected value 4 for key 4, got %d", got)
	}
}
