package main

import (
	"fmt"
	"mhy/utils"
)

//160.相交链表.intersection-of-two-linked-lists

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
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := getIntersectionNode(nil, nil)
	fmt.Println(result)
}

func getIntersectionNode_zwq(headA, headB *ListNode) *ListNode {
	//zwq
	pa, pb := headA, headB
	for pa != nil && pa.Next != nil {
		pa = pa.Next
	}
	for pb != nil && pb.Next != nil {
		pb = pb.Next
	}
	if pa != pb {
		return nil
	}
	//mine
	pa, pb = headA, headB
	for pa != nil && pb != nil {
		if pa == pb {
			return pa
		}
		pa = pa.Next
		pb = pb.Next
		if pa == nil {
			pa = headB
		}
		if pb == nil {
			pb = headA
		}
	}
	return nil
}

func getIntersectionNode_mine(headA, headB *ListNode) *ListNode {
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
