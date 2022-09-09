package main

import (
	"mhy/utils"
)

//92.反转链表 II.reverse-linked-list-ii

type ListNode = utils.SListNode

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	p, tq := head, head
	if left == 1 {
		p = dummy
	}
	for i := 0; i < left-2; i++ {
		p = p.Next
	}
	for i := 0; i < right-1; i++ {
		tq = tq.Next
	}
	//fmt.Println(p.Val)
	//fmt.Println(tq.Val)
	tp := p.Next
	q := tq.Next
	tq.Next = nil
	p.Next = reverseList(tp)
	tp.Next = q
	return dummy.Next
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// without t
	// head = head.Next.Next

	t := head.Next
	res := reverseList(head.Next)
	head.Next = nil
	t.Next = head
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	a := utils.CreatSingleList([]int{1, 2, 3})
	result := reverseBetween(a, 1, 3)
	utils.PrintList(result)

}
