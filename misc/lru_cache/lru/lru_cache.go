package lru

import "fmt"

// LinkedListNode struct defining a node in a linked list
type LinkedListNode[K comparable, V any] struct {
	key   K
	value V
	next  *LinkedListNode[K, V]
	back  *LinkedListNode[K, V]
}

// NewNode constructor for a new linked list node
func NewNode[K comparable, V any](key K, value V) *LinkedListNode[K, V] {
	return &LinkedListNode[K, V]{
		key:   key,
		value: value,
		next:  nil,
		back:  nil,
	}
}

// LruCache struct defining the LRU cache implementation
type LruCache[K comparable, V any] struct {
	size  int
	items map[K]*LinkedListNode[K, V]
	head  *LinkedListNode[K, V]
	tail  *LinkedListNode[K, V]
}

// NewLruCache constructor for a new LRU cache
func NewLruCache[K comparable, V any](size int) *LruCache[K, V] {
	return &LruCache[K, V]{
		size:  size,
		items: make(map[K]*LinkedListNode[K, V], size),
		head:  nil,
		tail:  nil,
	}
}

// Size returns the maximum size of the cache
func (c *LruCache[K, V]) Size() int {
	return c.size
}

// IsEmpty returns true if the cache is empty otherwise false
func (c *LruCache[K, V]) IsEmpty() bool {
	return len(c.items) == 0
}

// Put inserts the elment into the cache
// - If the key exists then updates the value and move the node to the front
// - If the key does not exist and the cache is full, then removes the last element and inserts the new one at the front
// - If the key does not exist and the cache is not full, then inserts the new element at the front
func (c *LruCache[K, V]) Put(key K, value V) {
	if node, exists := c.items[key]; exists {
		node.value = value
		c.moveToFront(node)
		return
	}

	newNode := NewNode(key, value)
	newNode.key = key

	if len(c.items) == c.size {
		delete(c.items, c.tail.key)
		c.removeNode(c.tail)
	}

	// Add the new node to the front
	c.addToFront(newNode)
	c.items[key] = newNode
}

// Get retrieves the value of a key and moves the corresponding node to the front of the list
func (c *LruCache[K, V]) Get(key K) (V, bool) {
	if node, ok := c.items[key]; ok {
		c.moveToFront(node)
		return node.value, true
	}

	var zeroValue V
	return zeroValue, false
}

// moveToFront moves a node to the front of the linked list
func (c *LruCache[K, V]) moveToFront(node *LinkedListNode[K, V]) {
	if node == c.head {
		return
	}

	c.removeNode(node)
	c.addToFront(node)
}

// addToFront adds a node to the front of the linked list
func (c *LruCache[K, V]) addToFront(node *LinkedListNode[K, V]) {
	node.next = c.head
	node.back = nil

	if c.head != nil {
		c.head.back = node
	}
	c.head = node

	if c.tail == nil {
		c.tail = node
	}
}

// removeNode removes a node from the linked list
func (c *LruCache[K, V]) removeNode(node *LinkedListNode[K, V]) {
	if node.back != nil {
		node.back.next = node.next
	} else {
		c.head = node.next
	}

	if node.next != nil {
		node.next.back = node.back
	} else {
		c.tail = node.back
	}
}

// String implements the Stsringer interface
func (c *LruCache[K, V]) String() string {
	var str string
	node := c.head
	for node != nil {
		str += fmt.Sprintf("%v: %v", node.key, node.value)
		if node.next != nil {
			str += fmt.Sprintf(" -> %v\n", node.next.key)
		}
		node = node.next
	}

	return str
}
