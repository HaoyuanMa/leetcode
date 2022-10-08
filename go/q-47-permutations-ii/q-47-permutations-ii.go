package main

import "fmt"

//47.全排列 II.permutations-ii

// leetcode submit region begin(Prohibit modification and deletion)
var ans [][]int
var tmp []int
var used []bool

func permuteUnique(nums []int) [][]int {
	ans = [][]int{}
	tmp = []int{}
	used = make([]bool, len(nums))
	backtrack(nums, 0)
	return ans
}

func backtrack(nums []int, cur int) {
	size := len(nums)
	if cur == size {
		tmpAns := make([]int, size)
		copy(tmpAns, tmp)
		ans = append(ans, tmpAns)
		return
	}
	usedV := make(map[int]bool, size)
	for i := 0; i < size; i++ {
		if usedV[nums[i]] || used[i] {
			continue
		}
		usedV[nums[i]] = true
		used[i] = true
		tmp = append(tmp, nums[i])
		backtrack(nums, cur+1)
		tmp = tmp[:len(tmp)-1]
		used[i] = false
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := permuteUnique([]int{1, 1, 2})
	fmt.Println(result)
}
