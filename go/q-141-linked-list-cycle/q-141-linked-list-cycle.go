package main

import (
	"fmt"
	"mhy/utils"
)

//141.环形链表.linked-list-cycle

type ListNode = utils.SListNode

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast, slow := head.Next, head
	for fast != nil && fast.Next != nil {
		if fast == slow {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := hasCycle(nil)
	fmt.Println(result)
}
