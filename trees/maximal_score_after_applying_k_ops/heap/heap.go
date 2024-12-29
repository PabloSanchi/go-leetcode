// Package heap
// Max heap implementation
package heap

type Heap struct {
	values []int
}

func NewHeap() *Heap {
	return &Heap{}
}

func (h *Heap) Size() int {
	return len(h.values)
}

func (h *Heap) Peek() (int, bool) {
	if len(h.values) == 0 {
		return 0, false
	}

	return h.values[0], true
}

func (h *Heap) Add(value int) {
	h.values = append(h.values, value)
	h.heapifyUp()
}

func (h *Heap) Delete() (int, bool) {
	if len(h.values) == 0 {
		return 0, false
	}

	value := h.values[0]
	h.values[0] = h.values[len(h.values)-1]
	h.values = h.values[:len(h.values)-1]
	h.heapifyDown()

	return value, true
}

// private

func (h *Heap) heapifyUp() {
	index := len(h.values) - 1
	for h.hasParent(index) && h.parent(index) < h.values[index] {
		h.swap(h.getParentIndex(index), index)
		index = h.getParentIndex(index)
	}
}

func (h *Heap) heapifyDown() {
	index := 0
	for h.hasLeftChild(index) {
		smallerChildIndex := h.getLeftChildIndex(index)
		if h.hasRightChild(index) && h.rightChild(index) > h.leftChild(index) {
			smallerChildIndex = h.getRightChildIndex(index)
		}

		if h.values[index] > h.values[smallerChildIndex] {
			break
		}

		h.swap(index, smallerChildIndex)
		index = smallerChildIndex
	}
}

func (h *Heap) getLeftChildIndex(parentIndex int) int {
	return 2*parentIndex + 1
}

func (h *Heap) getRightChildIndex(parentIndex int) int {
	return 2*parentIndex + 2
}

func (h *Heap) getParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

func (h *Heap) hasLeftChild(index int) bool {
	return h.getLeftChildIndex(index) < len(h.values)
}

func (h *Heap) hasRightChild(index int) bool {
	return h.getRightChildIndex(index) < len(h.values)
}

func (h *Heap) hasParent(index int) bool {
	return h.getParentIndex(index) >= 0
}

func (h *Heap) leftChild(index int) int {
	return h.values[h.getLeftChildIndex(index)]
}

func (h *Heap) rightChild(index int) int {
	return h.values[h.getRightChildIndex(index)]
}

func (h *Heap) parent(index int) int {
	return h.values[h.getParentIndex(index)]
}

func (h *Heap) swap(first, second int) {
	temp := h.values[first]
	h.values[first] = h.values[second]
	h.values[second] = temp
}
