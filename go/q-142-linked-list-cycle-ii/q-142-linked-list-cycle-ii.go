package main

import (
	"fmt"
	"mhy/utils"
)

//142.环形链表 II.linked-list-cycle-ii

type ListNode = utils.SListNode

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			fast = head
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	for {
		if fast == slow {
			return fast
		}
		fast = fast.Next
		slow = slow.Next

	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := detectCycle(nil)
	result = detectCycle_1(nil)
	fmt.Println(result)
}

func detectCycle_1(head *ListNode) *ListNode {
	mark := make(map[*ListNode]bool)
	for head != nil {
		_, ok := mark[head]
		if ok {
			return head
		}
		mark[head] = true
		head = head.Next
	}
	return nil
}
