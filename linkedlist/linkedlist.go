package linkedlist

import "fmt"

// ListNode linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// MakeLinkList make a linked list form int slice
func MakeLinkList(data []int) (head *ListNode) {
	head = new(ListNode)
	p := head
	for _, item := range data {
		p.Next = new(ListNode)
		p = p.Next
		p.Val = item
	}
	p.Next = nil
	head = head.Next
	return
}

// ListNodeEquals judge whether two linked list equals
func ListNodeEquals(l1, l2 *ListNode) bool {
	if l1 == nil || l2 == nil {
		return l1 == l2
	}
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			break
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == l2
}

func (l *ListNode) String() (result string) {

	for l != nil {
		result += fmt.Sprintf("%v-> ", l.Val)
		l = l.Next
	}
	result += "&"
	return
}
