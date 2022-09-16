package main

import "fmt"
import "mhy/utils"

//104.二叉树的最大深度.maximum-depth-of-binary-tree

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
var depth int
var ans int

func maxDepth(root *TreeNode) int {
	depth, ans = 0, 0
	traverse(root)
	return ans
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	depth++
	if root.Left == nil && root.Right == nil {
		if depth > ans {
			ans = depth
		}
	}
	traverse(root.Left)
	traverse(root.Right)
	depth--
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	root := &TreeNode{
		1,
		nil,
		&TreeNode{
			2,
			nil,
			nil,
		},
	}
	result := maxDepth(root)
	fmt.Println(result)
}

func maxDepth_recursion(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	if l > r {
		return l + 1
	} else {
		return r + 1
	}
}
