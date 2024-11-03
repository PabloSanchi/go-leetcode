package list

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered] struct {
	Val  T
	Next *Node[T]
	Prev *Node[T]
}

// newNode creates a new node
func newNode[T constraints.Ordered](value T) *Node[T] {
	return &Node[T]{Val: value}
}

type List[T constraints.Ordered] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

// NewList creates a new list
func NewList[T constraints.Ordered]() *List[T] {
	return &List[T]{}
}

// PushBack adds a node to the end of the list
func (l *List[T]) PushBack(value T) {
	node := newNode(value)

	if l.size == 0 {
		l.head = node
		l.tail = node
	} else {
		l.tail.Next = node
		node.Prev = l.tail
		l.tail = node
	}

	l.size++
}

// Push adds a node to the front of the list
func (l *List[T]) Push(value T) {
	node := newNode(value)

	if l.size == 0 {
		l.head = node
		l.tail = node
	} else {
		node.Next = l.head
		l.head.Prev = node
		l.head = node
	}

	l.size++
}

// PopBack removes a node from the end of the list
func (l *List[T]) PopBack() (*Node[T], bool) {
	if l.size == 0 {
		return nil, false
	}

	node := l.tail
	l.tail = l.tail.Prev
	if l.tail != nil {
		l.tail.Next = nil
	} else {
		l.head = nil // List is now empty
	}

	l.size--
	return node, true
}

// Pop removes the head node and returns its value
func (l *List[T]) Pop() (*Node[T], bool) {
	if l.size == 0 {
		return nil, false
	}

	node := l.head
	l.head = l.head.Next
	if l.head != nil {
		l.head.Prev = nil
	} else {
		l.tail = nil // List is now empty
	}

	l.size--
	return node, true
}

// At method returns the node at the given index
func (l *List[T]) At(index int) *Node[T] {
	if index < 0 || index >= l.size {
		return nil
	}

	node := l.head
	for i := 0; i < index; i++ {
		node = node.Next
	}

	return node
}

// Size returns the current size of the list
func (l *List[T]) Size() int {
	return l.size
}

func (l *List[T]) String() string {
	var str string
	node := l.head
	for node != nil {
		str += fmt.Sprintf("%v ", node.Val)
		node = node.Next
	}

	return str
}
