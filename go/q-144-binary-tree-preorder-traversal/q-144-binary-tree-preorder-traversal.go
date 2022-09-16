package main

import (
	"fmt"
	"mhy/utils"
)

//144.二叉树的前序遍历.binary-tree-preorder-traversal

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

var fans []int

func preorderTraversal(root *TreeNode) []int {
	fans = []int{}
	traverse(root)
	return fans
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	fans = append(fans, root.Val)
	traverse(root.Left)
	traverse(root.Right)
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := preorderTraversal(nil)
	fmt.Println(result)
}

func preorderTraversal_recursion(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	ans = append(ans, root.Val)
	l := preorderTraversal(root.Left)
	r := preorderTraversal(root.Right)
	ans = append(ans, l...)
	ans = append(ans, r...)
	return ans
}
