package main

import "fmt"

//21.合并两个有序链表.merge-two-sorted-lists

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var ans *ListNode
	if list1.Val < list2.Val {
		ans = list1
		list1 = list1.Next
	} else {
		ans = list2
		list2 = list2.Next
	}

	cur := ans
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}

	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1 *ListNode = nil
	var l2 *ListNode = nil
	var result = mergeTwoLists(l1, l2)
	fmt.Println(result)
}
