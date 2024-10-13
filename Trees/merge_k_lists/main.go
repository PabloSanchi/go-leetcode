package main

import (
	"fmt"
	"mergeklists/heap"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(value int) *ListNode {
	return &ListNode{
		Val:  value,
		Next: nil,
	}
}

// ToSlice will convert the linked list to a slice of integers (only for testing)
func (l *ListNode) ToSlice() []int {
	result := []int{}
	head := l
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}

	return result
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := heap.NewHeap()

	for _, list := range lists {
		head := list
		for head != nil {
			h.Add(head.Val)
			head = head.Next
		}
	}

	result := NewListNode(0)
	dummyHead := result
	for {
		value, ok := h.Delete()
		if !ok {
			break
		}

		result.Next = NewListNode(value)
		result = result.Next
	}

	return dummyHead.Next
}

func main() {
	// Input: lists = [[1,4,5],[1,3,4],[2,6]]
	// Output: [1,1,2,3,4,4,5,6]
	head1 := NewListNode(1)
	head1.Next = NewListNode(4)
	head1.Next.Next = NewListNode(5)

	head2 := NewListNode(1)
	head2.Next = NewListNode(3)
	head2.Next.Next = NewListNode(4)

	head3 := NewListNode(2)
	head3.Next = NewListNode(6)

	result := mergeKLists([]*ListNode{head1, head2, head3})

	fmt.Printf("Result ")
	for result != nil {
		fmt.Printf("-> %d ", result.Val)
		result = result.Next
	}
}
