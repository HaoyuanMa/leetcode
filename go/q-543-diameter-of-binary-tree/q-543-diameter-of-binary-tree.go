package main

import (
	"fmt"
	"mhy/utils"
)

//543.二叉树的直径.diameter-of-binary-tree

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

func diameterOfBinaryTree(root *TreeNode) int {
	ans, _ := solve(root)
	return ans
}

// 兼顾速度与空间
func solve(root *TreeNode) (ans, depth int) {
	if root == nil {
		return 0, 0
	}
	if root.Left == nil && root.Right == nil {
		return 0, 1
	}
	leftAns, leftDepth := solve(root.Left)
	rightAns, rightDepth := solve(root.Right)

	ans = leftAns
	if rightAns > ans {
		ans = rightAns
	}

	tmpAns := leftDepth + rightDepth
	if tmpAns > ans {
		ans = tmpAns
	}

	tmpDepth := leftDepth
	if rightDepth > tmpDepth {
		tmpDepth = rightDepth
	}
	depth = tmpDepth + 1
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := diameterOfBinaryTree(nil)
	fmt.Println(result)
}

// map优化
// 速度快但需要额外的空间
var depth = make(map[*TreeNode]int)

func diameterOfBinaryTree_map(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 0
	}
	l := diameterOfBinaryTree(root.Left)
	r := diameterOfBinaryTree(root.Right)
	ans := 0
	if l > r {
		ans = l
	} else {
		ans = r
	}
	t := maxDepth(root.Left) + maxDepth(root.Right)
	if t > ans {
		ans = t
	}
	return ans
}

func maxDepth(root *TreeNode) int {
	d, ok := depth[root]
	if ok {
		return d
	}
	if root == nil {
		depth[root] = 0
		return 0
	}
	if root.Left == nil && root.Right == nil {
		depth[root] = 1
		return 1
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	if l > r {
		depth[root] = l + 1
		return l + 1
	} else {
		depth[root] = r + 1
		return r + 1
	}
}
