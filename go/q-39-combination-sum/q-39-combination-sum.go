package main

import "fmt"

//39.组合总和.combination-sum

// leetcode submit region begin(Prohibit modification and deletion)
var ans [][]int
var tmp []int
var sum int

func combinationSum(candidates []int, target int) [][]int {
	ans = [][]int{}
	tmp = []int{}
	sum = 0
	backtrack(candidates, target, 0)
	return ans
}

func backtrack(candidates []int, target, start int) {
	if sum > target {
		return
	}
	size := len(candidates)
	if sum == target {
		tmpAns := make([]int, len(tmp))
		copy(tmpAns, tmp)
		ans = append(ans, tmpAns)
	}
	for i := start; i < size; i++ {
		tmp = append(tmp, candidates[i])
		sum += candidates[i]
		backtrack(candidates, target, i)
		sum -= candidates[i]
		tmp = tmp[:len(tmp)-1]
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := combinationSum([]int{2, 3, 6, 7}, 7)
	fmt.Println(result)
}
