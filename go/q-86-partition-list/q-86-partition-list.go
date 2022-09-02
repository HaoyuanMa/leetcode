package main

import "fmt"

//86.分隔链表.partition-list

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	p, q := &ListNode{}, &ListNode{}
	l, r := p, q
	for head != nil {
		if head.Val < x {
			p.Next = head
			p = p.Next
			head = head.Next
			p.Next = nil
		} else {
			q.Next = head
			q = q.Next
			head = head.Next
			q.Next = nil
		}
	}
	p.Next = r.Next
	return l.Next
}

//leetcode submit region end(Prohibit modification and deletion)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var x int = 3
	head := &ListNode{
		Val:  1,
		Next: nil,
	}
	result := partition(head, x)
	fmt.Println(result)
}
