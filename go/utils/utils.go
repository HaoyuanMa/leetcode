package utils

import "fmt"

type SListNode struct {
	Val  int
	Next *SListNode
}

func CreatSingleList(data []int) *SListNode {
	head := &SListNode{}
	l := len(data)
	cur := head
	for i := 0; i < l; i++ {
		cur.Next = &SListNode{
			Val:  data[i],
			Next: nil,
		}
		cur = cur.Next
	}
	return head.Next
}

func PrintList(list *SListNode) {
	cur := list
	for cur != nil {
		fmt.Print(cur.Val)
		fmt.Print("  ")
		cur = cur.Next
	}
}
