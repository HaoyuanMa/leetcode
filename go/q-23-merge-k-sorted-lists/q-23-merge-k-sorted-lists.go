package main

import (
	"mhy/utils"
)

//23.合并K个升序链表.merge-k-sorted-lists

type ListNode = utils.ListNode

//leetcode submit region begin(Prohibit modification and deletion)
/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)
	if l == 0 {
		return nil
	} else if l == 1 {
		return lists[0]
	} else if l == 2 {
		left := lists[0]
		right := lists[1]
		head := &ListNode{}
		cur := head
		for left != nil && right != nil {
			if left.Val < right.Val {
				cur.Next = left
				left = left.Next
			} else {
				cur.Next = right
				right = right.Next
			}
			cur = cur.Next
		}
		if left != nil {
			cur.Next = left
		}
		if right != nil {
			cur.Next = right
		}
		return head.Next
	} else {
		// l > 2
		mid := l / 2
		left := mergeKLists(lists[:mid])
		right := mergeKLists(lists[mid:])
		return mergeKLists([]*ListNode{left, right})
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

	a := utils.CreatSingleList([]int{1, 4, 5})
	b := utils.CreatSingleList([]int{1, 3, 4})
	c := utils.CreatSingleList([]int{2, 6})

	result := mergeKLists([]*ListNode{a, b, c})
	utils.PrintList(result)
}
