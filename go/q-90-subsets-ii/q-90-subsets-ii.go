package main

import (
	"fmt"
	"sort"
)

//90.子集 II.subsets-ii

//leetcode submit region begin(Prohibit modification and deletion)

var ans [][]int
var tmp []int

func subsetsWithDup(nums []int) [][]int {
	ans = [][]int{}
	tmp = []int{}
	sort.Ints(nums)
	backtrack(nums, 0)
	return ans
}

func backtrack(nums []int, start int) {
	tmpAns := make([]int, len(tmp))
	copy(tmpAns, tmp)
	ans = append(ans, tmpAns)
	size := len(nums)
	//used := map[int]bool{}
	for i := start; i < size; i++ {
		//if used[nums[i]] {
		//	continue
		//}
		//used[nums[i]] = true
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		tmp = append(tmp, nums[i])
		backtrack(nums, i+1)
		tmp = tmp[:len(tmp)-1]
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := subsetsWithDup([]int{1, 2, 1})
	fmt.Println(result)
}
