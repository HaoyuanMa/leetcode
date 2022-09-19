package main

import (
	"fmt"
	"mhy/utils"
)

//94.二叉树的中序遍历.binary-tree-inorder-traversal

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

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ans []int
	ans = append(ans, inorderTraversal(root.Left)...)
	ans = append(ans, root.Val)
	ans = append(ans, inorderTraversal(root.Right)...)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := inorderTraversal(nil)
	fmt.Println(result)
}
