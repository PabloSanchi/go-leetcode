package main

import (
	"reflect"
	"testing"
)

// Input: lists = [[1,4,5],[1,3,4],[2,6]]
// Output: [1,1,2,3,4,4,5,6]
func TestCase1(t *testing.T) {
	head1 := NewListNode(1)
	head1.Next = NewListNode(4)
	head1.Next.Next = NewListNode(5)

	head2 := NewListNode(1)
	head2.Next = NewListNode(3)
	head2.Next.Next = NewListNode(4)

	head3 := NewListNode(2)
	head3.Next = NewListNode(6)

	result := mergeKLists([]*ListNode{head1, head2, head3})

	expected := []int{1, 1, 2, 3, 4, 4, 5, 6}

	if !reflect.DeepEqual(result.ToSlice(), expected) {
		t.Errorf("Expected %v but got %v", expected, result.ToSlice())
	}
}

// Input: lists = []
// Output: []
func TestCase2(t *testing.T) {
	result := mergeKLists([]*ListNode{})

	expected := []int{}

	if !reflect.DeepEqual(result.ToSlice(), expected) {
		t.Errorf("Expected %v but got %v", expected, result.ToSlice())
	}
}

// Input: lists = [[]]
// Output: []
func TestCase3(t *testing.T) {
	result := mergeKLists([]*ListNode{nil})

	expected := []int{}

	if !reflect.DeepEqual(result.ToSlice(), expected) {
		t.Errorf("Expected %v but got %v", expected, result.ToSlice())
	}
}
