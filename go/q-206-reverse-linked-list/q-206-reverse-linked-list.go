package main

import (
	"mhy/utils"
)

//206.反转链表.reverse-linked-list

type ListNode = utils.SListNode

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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
	result := reverseList(a)
	utils.PrintList(result)
}

func reverseList_recursion(head *ListNode) *ListNode {
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
