package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// RotateRight rotates the list to the right by k places.
// https://leetcode.com/problems/rotate-list/
// Overall explanation:
// 1 -> 2 -> 3 -> 4 -> 5
// n = 5, k = 2; n-k = 3
// newTail = 3
// newTail.Next = 4 so the new head is 4
// the tail shoud point to nil, so newTail.Next = nil
// the old tail is 5, and now points to the initial head 1
// done
func RotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// get the length of the list
	n := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		n++
	}

	// how many time to rotate to avoid repeated rotations
	// if we have a list of 5 and rotate 10 times it's the same as rotating 5 times
	// and rotating 5 times is the same as not rotating at all
	k %= n
	if k == 0 {
		return head
	}

	// find the node that will be the new tail
	newTail := head
	for i := 1; i < n-k; i++ {
		newTail = newTail.Next
	}

	// get the next node that will be the new head
	newHead := newTail.Next
	// as is a tail, the next node is nil
	newTail.Next = nil
	// make the last node of the first list point to the initial head
	tail.Next = head

	return newHead
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	head = RotateRight(head, 3)

	for head != nil {
		println(head.Val)
		head = head.Next
	}
}
