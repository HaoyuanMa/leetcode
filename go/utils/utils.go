package utils

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreatSingleList(data []int) *ListNode {
	head := &ListNode{}
	l := len(data)
	cur := head
	for i := 0; i < l; i++ {
		cur.Next = &ListNode{
			Val:  data[i],
			Next: nil,
		}
		cur = cur.Next
	}
	return head.Next
}

func PrintList(list *ListNode) {
	cur := list
	for cur != nil {
		fmt.Print(cur.Val)
		fmt.Print("  ")
		cur = cur.Next
	}
}
