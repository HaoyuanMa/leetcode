package main

import (
	"fmt"
)

//51.N 皇后.n-queens

// leetcode submit region begin(Prohibit modification and deletion)
var (
	used [][]bool
	cur  []int
	ans  [][]string
)

func solveNQueens(n int) [][]string {
	cur = []int{}
	ans = [][]string{}
	used = make([][]bool, n)
	for i := 0; i < n; i++ {
		used[i] = make([]bool, n)
	}
	backtrack(0, n)
	return ans
}

func backtrack(row, n int) {
	if row == n {
		var solution []string
		for i := 0; i < n; i++ {
			line := ""
			for j := 0; j < n; j++ {
				if cur[i] == j {
					line += "Q"
				} else {
					line += "."
				}
			}
			solution = append(solution, line)
		}
		ans = append(ans, solution)
		return
	}
	for i := 0; i < n; i++ {
		if hasCollision(i, row, n, used) {
			continue
		}
		cur = append(cur, i)
		used[row][i] = true
		backtrack(row+1, n)
		cur = cur[:len(cur)-1]
		used[row][i] = false
	}
}
func hasCollision(col, row, n int, used [][]bool) (collision bool) {
	collision = false
	//col
	for j := 0; j < n; j++ {
		if j == col {
			continue
		}
		if used[row][j] {
			collision = true
			break
		}
	}
	//row
	for j := 0; j < n; j++ {
		if j == row {
			continue
		}
		if used[j][col] {
			collision = true
			break
		}
	}
	//down-right
	for j := 0; j < n; j++ {
		if row+j == n || col+j == n {
			break
		}
		if used[row+j][col+j] {
			collision = true
			break
		}
	}
	//down-left
	for j := 0; j < n; j++ {
		if row+j == n || col-j < 0 {
			break
		}
		if used[row+j][col-j] {
			collision = true
			break
		}
	}
	//up-left
	for j := 0; j < n; j++ {
		if row-j < 0 || col-j < 0 {
			break
		}
		if used[row-j][col-j] {
			collision = true
			break
		}
	}
	//up-right
	for j := 0; j < n; j++ {
		if row-j < 0 || col+j == n {
			break
		}
		if used[row-j][col+j] {
			collision = true
			break
		}
	}
	return collision
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := solveNQueens(4)
	fmt.Println(result)
}
