package main

import (
	"fmt"
	"mhy/utils"
)

//145.二叉树的后序遍历.binary-tree-postorder-traversal

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
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ans := postorderTraversal(root.Left)
	ans = append(ans, postorderTraversal(root.Right)...)
	ans = append(ans, root.Val)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := postorderTraversal(nil)
	fmt.Println(result)
}
