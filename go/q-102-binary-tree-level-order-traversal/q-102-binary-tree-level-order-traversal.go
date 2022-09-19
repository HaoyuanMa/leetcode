package main

import (
	"fmt"
	"mhy/utils"
)

//102.二叉树的层序遍历.binary-tree-level-order-traversal

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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ans [][]int
	queue := []*TreeNode{root}

	for size := len(queue); size > 0; size = len(queue) {
		var cur []int
		for i := 0; i < size; i++ {
			if queue[i] == nil {
				continue
			}
			cur = append(cur, queue[i].Val)
			queue = append(queue, queue[i].Left)
			queue = append(queue, queue[i].Right)
		}
		if cur != nil {
			ans = append(ans, cur)
		}
		queue = queue[size:]
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	t := &TreeNode{
		0,
		&TreeNode{
			1,
			&TreeNode{3, nil, &TreeNode{5, nil, nil}},
			nil,
		},
		&TreeNode{
			2,
			&TreeNode{4, nil, nil},
			nil,
		},
	}
	result := levelOrder(t)
	fmt.Println(result)
}
