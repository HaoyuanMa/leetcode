package main

import (
	"mhy/utils"
)

//83.删除排序链表中的重复元素.remove-duplicates-from-sorted-list

type ListNode = utils.SListNode

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p, q := head, head.Next
	for q != nil {
		if p.Val != q.Val {
			p.Next = q
			p = q
		}
		q = q.Next
	}
	p.Next = nil
	return head
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	a := []int{1, 2, 2, 3, 4, 4, 5, 6, 6}
	result := deleteDuplicates(utils.CreatSingleList(a))
	utils.PrintList(result)
}
