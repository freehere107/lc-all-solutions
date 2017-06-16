package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{3, nil}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{4, nil}}}
	fmt.Print(addTwoNumbers(l1, l2))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	prev := &ListNode{}
	carry := 0
	result := prev
	for {
		if l1 == nil || l2 == nil {
			break
		} else {
			data := l1.Val + l2.Val + carry
			l1 = l1.Next
			l2 = l2.Next
			val := data % 10
			carry = data / 10
			prev.Next = &ListNode{Val: val, Next: nil}
			prev = prev.Next
		}
	}
	return result.Next
}
