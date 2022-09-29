package main

import (
	"fmt"
	"sort"
)

//40.组合总和 II.combination-sum-ii

//leetcode submit region begin(Prohibit modification and deletion)

var ans [][]int
var tmp []int
var sum int

func combinationSum2(candidates []int, target int) [][]int {
	ans = [][]int{}
	tmp = []int{}
	sum = 0
	sort.Ints(candidates)
	backtrack(0, target, candidates)
	return ans
}

func backtrack(start, target int, nums []int) {
	if sum > target {
		return
	}
	if sum == target {
		tmpAns := make([]int, len(tmp))
		copy(tmpAns, tmp)
		ans = append(ans, tmpAns)
	}

	size := len(nums)
	for i := start; i < size; i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}

		tmp = append(tmp, nums[i])
		sum += nums[i]
		backtrack(i+1, target, nums)
		sum -= nums[i]
		tmp = tmp[:len(tmp)-1]
	}

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	fmt.Println(result)
}
