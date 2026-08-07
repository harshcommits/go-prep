package syntaxtest

import "container/list"

// LRUCache implements a least-recently-used cache using a doubly linked list
// and a map for O(1) lookups.
type LRUCache struct {
	cap  int
	list *list.List
	m    map[int]int
}

// NewLRUCache creates a new LRU cache with the given capacity.
func NewLRUCache(capacity int) *LRUCache {
	if capacity <= 0 {
		return &LRUCache{cap: 0, list: list.New(), m: make(map[int]int)}
	}

	return &LRUCache{
		cap:  capacity,
		list: list.New(),
		m:    make(map[int]int, capacity),
	}
}

// Get returns the value for a key, or -1 if the key is not present.
// It also marks the key as recently used.
func (c *LRUCache) Get(key int) int {
	if c == nil {
		return -1
	}
	if value, ok := c.m[key]; ok {
		c.moveToFront(key)
		return value
	}
	return -1
}

// Put adds or updates a key-value pair in the cache.
// If the cache is at capacity, the least recently used item is evicted.
func (c *LRUCache) Put(key, value int) {
	if c == nil {
		return
	}

	if _, ok := c.m[key]; ok {
		c.m[key] = value
		c.moveToFront(key)
		return
	}

	if c.cap > 0 && len(c.m) >= c.cap {
		c.removeLeastRecentlyUsed()
	}

	c.m[key] = value
	c.list.PushFront(key)
}

func (c *LRUCache) moveToFront(key int) {
	for e := c.list.Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			c.list.MoveToFront(e)
			return
		}
	}
}

func (c *LRUCache) removeLeastRecentlyUsed() {
	if c.list.Len() == 0 {
		return
	}

	back := c.list.Back()
	if back == nil {
		return
	}

	key := back.Value.(int)
	delete(c.m, key)
	c.list.Remove(back)
}
