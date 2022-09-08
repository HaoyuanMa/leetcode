package main

import (
	"mhy/utils"
)

//876.链表的中间结点.middle-of-the-linked-list

type ListNode = utils.SListNode

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	front, back := head, head
	//for front != nil {
	//	front = front.Next
	//	if front != nil {
	//		back = back.Next
	//		front = front.Next
	//	}
	//}
	for front != nil && front.Next != nil {
		front = front.Next.Next
		back = back.Next
	}
	return back
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := middleNode(utils.CreatSingleList(a))
	utils.PrintList(result)
}
