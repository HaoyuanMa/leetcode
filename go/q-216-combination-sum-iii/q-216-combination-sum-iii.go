package main

import "fmt"

//216.组合总和 III.combination-sum-iii

// leetcode submit region begin(Prohibit modification and deletion)
var ans [][]int
var tmp []int
var sum int

func combinationSum3(k int, n int) [][]int {
	ans = [][]int{}
	tmp = []int{}
	sum = 0
	backtrack(k, n, 1)
	return ans
}

func backtrack(k, n, start int) {
	if sum > n {
		return
	}
	if sum == n && len(tmp) == k {
		tmpAns := make([]int, k)
		copy(tmpAns, tmp)
		ans = append(ans, tmpAns)
	}
	for i := start; i < 10; i++ {
		tmp = append(tmp, i)
		sum += i
		backtrack(k, n, i+1)
		sum -= i
		tmp = tmp[:len(tmp)-1]
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := combinationSum3(3, 7)
	fmt.Println(result)
}
