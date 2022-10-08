package main

import "fmt"

//77.组合.combinations

// leetcode submit region begin(Prohibit modification and deletion)
var ans [][]int
var tmp []int

func combine(n int, k int) [][]int {
	ans = [][]int{}
	tmp = []int{}
	backtrack(n, k, 1)
	return ans
}

func backtrack(n, k, start int) {
	if len(tmp) == k {
		tmpAns := make([]int, k)
		copy(tmpAns, tmp)
		ans = append(ans, tmpAns)
		return
	}
	for i := start; i <= n; i++ {
		tmp = append(tmp, i)
		backtrack(n, k, i+1)
		tmp = tmp[:len(tmp)-1]
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := combine(4, 2)
	fmt.Println(result)
}
