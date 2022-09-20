package main

import "fmt"

//46.全排列.permutations

// leetcode submit region begin(Prohibit modification and deletion)

var ans [][]int
var cur []int
var used [10]bool

func permute(nums []int) [][]int {
	ans = [][]int{}
	cur = []int{}
	used = [10]bool{}
	backtrack(0, nums)
	return ans
}

func backtrack(depth int, nums []int) {
	if depth == len(nums) {
		var newAns []int
		newAns = append(newAns, cur...)
		ans = append(ans, newAns)

		// ans = append(ans, cur)
		// 最终结果只保留了最后一次，如输入[1，2，3]输入
		// [[3,2,1],[3,2,1],[3,2,1],[3,2,1],[3,2,1],[3,2,1]]
		// cur地址不变，ans中存其地址，cur一变，ans也变

		return
	}
	for index, num := range nums {
		if used[index] {
			continue
		}
		cur = append(cur, num)
		used[index] = true
		backtrack(depth+1, nums)
		used[index] = false
		cur = cur[:len(cur)-1]
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := permute([]int{1, 2, 3})
	fmt.Println(result)
}
