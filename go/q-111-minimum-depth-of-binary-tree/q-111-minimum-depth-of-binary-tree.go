package main

import (
	"fmt"
	"mhy/utils"
)

//111.二叉树的最小深度.minimum-depth-of-binary-tree

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
var queue []*TreeNode

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue = []*TreeNode{}
	queue = append(queue, root)
	ans := 0

	for len(queue) > 0 {
		ans++
		curSize := len(queue)
		for j := 0; j < curSize; j++ {
			if queue[j].Left == nil && queue[j].Right == nil {
				return ans
			}
			if queue[j].Left != nil {
				queue = append(queue, queue[j].Left)
			}
			if queue[j].Right != nil {
				queue = append(queue, queue[j].Right)
			}
		}
		if len(queue) > curSize {
			queue = queue[curSize:]
		} else {
			queue = nil
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := minDepth(&TreeNode{1,
		&TreeNode{2,
			&TreeNode{4, nil, nil}, nil},
		&TreeNode{3,
			nil,
			&TreeNode{5, nil, nil}}})
	fmt.Println(result)
}
