package main

import (
	"fmt"
	"mhy/utils"
)

//337.打家劫舍 III.house-robber-iii

type TreeNode = utils.BiTreeNode

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(recursion(root))
}
func recursion(t *TreeNode) (notDo int, do int) {
	if t == nil {
		return 0, 0
	}
	l0, l1 := recursion(t.Left)
	r0, r1 := recursion(t.Right)
	notDo = max(l0, l1) + max(r0, r1)
	do = l0 + r0 + t.Val
	return
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// leetcode submit region end(Prohibit modification and deletion)
// mine begin
var caled0, caled1 map[*TreeNode]int

func rob_mine(root *TreeNode) int {
	if root == nil {
		return 0
	}
	caled0, caled1 = map[*TreeNode]int{}, map[*TreeNode]int{}
	return max(recursion0(root), recursion1(root))
}
func recursion0(t *TreeNode) int {
	if t == nil {
		return 0
	}
	if v, ok := caled0[t]; ok {
		return v
	}
	a0 := recursion0(t.Right) + recursion0(t.Left)
	a1 := recursion1(t.Right) + recursion1(t.Left)
	l := recursion1(t.Left) + recursion0(t.Right)
	r := recursion1(t.Right) + recursion0(t.Left)
	allMax := max(a0, a1)
	singleMax := max(l, r)
	caled0[t] = max(allMax, singleMax)
	return caled0[t]
}
func recursion1(t *TreeNode) int {
	if t == nil {
		return 0
	}
	if v, ok := caled1[t]; ok {
		return v
	}
	son := recursion0(t.Right) + recursion0(t.Left)
	caled1[t] = t.Val + son
	return caled1[t]
}

// mine end

// lbld begin
var caled = map[*TreeNode]int{}

func rob_lbld(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if v, ok := caled[root]; ok {
		return v
	}
	l, r := 0, 0
	if root.Left != nil {
		l = rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		r = rob(root.Right.Left) + rob(root.Right.Right)
	}
	s1 := root.Val + l + r
	s0 := rob(root.Left) + rob(root.Right)
	caled[root] = max(s0, s1)
	return caled[root]
}

//lbld end

func main() {
	result := rob(nil)
	fmt.Println(result)
}
