package main

import "fmt"

//78.子集.subsets

// leetcode submit region begin(Prohibit modification and deletion)

var tmp []int
var ans [][]int

func subsets(nums []int) [][]int {
	tmp = []int{}
	ans = [][]int{}
	backtrack(nums, 0)
	return ans
}

func backtrack(nums []int, start int) {
	tmpAns := make([]int, len(tmp))
	copy(tmpAns, tmp)
	ans = append(ans, tmpAns)
	size := len(nums)
	for i := start; i < size; i++ {
		tmp = append(tmp, nums[i])
		backtrack(nums, i+1)
		tmp = tmp[:len(tmp)-1]
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := subsets([]int{1, 2, 3, 4})
	fmt.Println(result)
}
