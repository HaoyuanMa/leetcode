package main

import (
	"fmt"
)

//304.二维区域和检索 - 矩阵不可变.range-sum-query-2d-immutable

// leetcode submit region begin(Prohibit modification and deletion)
type NumMatrix struct {
	matrix [][]int
	sum    [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := make([][]int, len(matrix))
	s := make([][]int, len(matrix))
	for i := 0; i < len(m); i++ {
		m[i] = make([]int, len(matrix[0]))
		copy(m[i], matrix[i])
		s[i] = make([]int, len(matrix[0]))
	}
	s[0][0] = m[0][0]
	for i := 1; i < len(m[0]); i++ {
		s[0][i] = s[0][i-1] + m[0][i]
	}
	for i := 1; i < len(m); i++ {
		s[i][0] = s[i-1][0] + m[i][0]
	}
	for i := 1; i < len(m); i++ {
		for j := 1; j < len(m[0]); j++ {
			s[i][j] = s[i][j-1] + s[i-1][j] + m[i][j] - s[i-1][j-1]
		}
	}
	return NumMatrix{m, s}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	ans := this.sum[row2][col2]
	if row1 > 0 {
		ans -= this.sum[row1-1][col2]
	}
	if col1 > 0 {
		ans -= this.sum[row2][col1-1]
	}
	if row1 > 0 && col1 > 0 {
		ans += this.sum[row1-1][col1-1]
	}
	return ans
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	mt := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	// [][]int{{1}, {2}, {3}, {4}}
	// [][]int{1,2,3,4}
	m := Constructor(mt)
	result := m.SumRegion(0, 0, 0, 0)
	fmt.Println(result)
}
