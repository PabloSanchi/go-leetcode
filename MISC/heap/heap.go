package heap

import (
	"golang.org/x/exp/constraints"
	"heap/list"
)

type Heap[T constraints.Ordered] struct {
	values list.List[T]
}

func NewHeap[T constraints.Ordered]() *Heap[T] {
	return &Heap[T]{
		values: *list.NewList[T](),
	}
}

func (h *Heap[T]) Size() int {
	return h.values.Size()
}

func (h *Heap[T]) Peek() (T, bool) {
	if node := h.values.At(0); node != nil {
		return node.Val, true
	}

	var zero T
	return zero, false
}

func (h *Heap[T]) Add(value T) {
	h.values.PushBack(value)
	h.heapifyUp()
}

func (h *Heap[T]) Delete() (T, bool) {
	if h.values.Size() == 0 {
		var zero T
		return zero, false
	}

	node := h.values.At(0)
	h.values.Pop()
	h.heapifyDown()

	return node.Val, true
}

func (h *Heap[T]) String() string {
	return h.values.String()
}

func (h *Heap[T]) heapifyUp() {
	index := h.values.Size() - 1
	for h.hasParent(index) && h.parent(index) > h.values.At(index).Val {
		h.swap(h.getParentIndex(index), index)
		index = h.getParentIndex(index)
	}
}

func (h *Heap[T]) heapifyDown() {
	index := 0
	for h.hasLeftChild(index) {
		smallerChildIndex := h.getLeftChildIndex(index)
		if h.hasRightChild(index) && h.rightChild(index) < h.leftChild(index) {
			smallerChildIndex = h.getRightChildIndex(index)
		}

		if h.values.At(index).Val < h.values.At(smallerChildIndex).Val {
			break
		}

		h.swap(index, smallerChildIndex)
		index = smallerChildIndex
	}
}

func (h *Heap[T]) getLeftChildIndex(parentIndex int) int {
	return 2*parentIndex + 1
}

func (h *Heap[T]) getRightChildIndex(parentIndex int) int {
	return 2*parentIndex + 2
}

func (h *Heap[T]) getParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

func (h *Heap[T]) hasLeftChild(index int) bool {
	return h.getLeftChildIndex(index) < h.values.Size()
}

func (h *Heap[T]) hasRightChild(index int) bool {
	return h.getRightChildIndex(index) < h.values.Size()
}

func (h *Heap[T]) hasParent(index int) bool {
	return h.getParentIndex(index) >= 0
}

func (h *Heap[T]) leftChild(index int) T {
	if node := h.values.At(h.getLeftChildIndex(index)); node != nil {
		return node.Val
	}

	var zero T
	return zero
}

func (h *Heap[T]) rightChild(index int) T {
	if node := h.values.At(h.getRightChildIndex(index)); node != nil {
		return node.Val
	}

	var zero T
	return zero
}

func (h *Heap[T]) parent(index int) T {
	if node := h.values.At(h.getParentIndex(index)); node != nil {
		return node.Val
	}

	var zero T
	return zero
}

func (h *Heap[T]) swap(first, second int) {
	firstNode := h.values.At(first)
	secondNode := h.values.At(second)
	firstNode.Val, secondNode.Val = secondNode.Val, firstNode.Val
}
