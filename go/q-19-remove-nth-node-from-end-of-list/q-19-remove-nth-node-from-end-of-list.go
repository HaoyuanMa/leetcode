package main

import (
	"mhy/utils"
)

//19.删除链表的倒数第 N 个结点.remove-nth-node-from-end-of-list

type ListNode = utils.SListNode

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	cur := dummy
	for i := 0; i <= n; i++ {
		cur = cur.Next
	}
	dest := dummy
	for cur != nil {
		cur = cur.Next
		dest = dest.Next
	}
	dest.Next = dest.Next.Next
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	a := []int{1, 2, 3, 4, 5}
	result := removeNthFromEnd(utils.CreatSingleList(a), 2)
	utils.PrintList(result)
}
