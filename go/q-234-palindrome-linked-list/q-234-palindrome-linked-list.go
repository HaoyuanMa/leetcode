package main

import (
	"fmt"
	"mhy/utils"
)

//234.回文链表.palindrome-linked-list

type ListNode = utils.SListNode

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	if fast != nil {
		slow = slow.Next
	} else {

	}
	p, q := head, reverseList(slow)

	for q != nil {
		if p.Val != q.Val {
			return false
		}
		p, q = p.Next, q.Next
	}
	return true
}

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
	a := utils.CreatSingleList([]int{1, 2, 3, 4, 3, 2, 1})
	result := isPalindrome(a)
	fmt.Println(result)
}

var left *ListNode

func isPalindrome_recursion(head *ListNode) bool {
	left = head
	return isValid(head)
}

func isValid(right *ListNode) bool {
	if right == nil {
		return true
	}
	t := isValid(right.Next)
	if t && left.Val == right.Val {
		left = left.Next
		return true
	}
	return false
}
