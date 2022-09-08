package main

import (
	"fmt"
	"mhy/utils"
)

//剑指 Offer 52.两个链表的第一个公共节点.liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof

type ListNode = utils.SListNode

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pa, pb := headA, headB
	for pa != nil && pb != nil {
		pa = pa.Next
		pb = pb.Next
	}
	var cur, curL, curS *ListNode
	if pa == nil {
		cur, curL, curS = pb, headB, headA
	}
	if pb == nil {
		cur, curL, curS = pa, headA, headB
	}
	for cur != nil {
		cur = cur.Next
		curL = curL.Next
	}
	for curL != nil {
		if curL == curS {
			return curL
		}
		curL = curL.Next
		curS = curS.Next
	}
	return curL
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := getIntersectionNode(nil, nil)
	fmt.Println(result)
}
