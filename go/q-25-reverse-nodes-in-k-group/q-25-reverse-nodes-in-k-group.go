package main

import (
	"mhy/utils"
)

//25.K 个一组翻转链表.reverse-nodes-in-k-group

type ListNode = utils.SListNode

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head
	var count int
	for count = 0; count < k-1 && p.Next != nil; count++ {
		p = p.Next
	}
	if count < k-1 {
		return head
	}
	rest := reverseKGroup(p.Next, k)
	p.Next = nil
	res := reverseList(head)
	head.Next = rest
	return res
}

func reverseList(head *ListNode) *ListNode {
	var dummy *ListNode = nil
	for head != nil {
		t := head.Next
		head.Next = dummy
		dummy = head
		head = t
	}
	return dummy
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	a := utils.CreatSingleList([]int{1, 2, 3, 4, 5})
	result := reverseKGroup(a, 1)
	utils.PrintList(result)
}
