package main

import "fmt"

//2379.得到 K 个黑块的最少涂色次数.minimum-recolors-to-get-k-consecutive-black-blocks

// leetcode submit region begin(Prohibit modification and deletion)
func minimumRecolors(blocks string, k int) int {
	var ans int = k
	var count int = 0
	var l int = len(blocks)
	for i := 0; i < k; i++ {
		if blocks[i] == 'W' {
			count++
		}
	}
	if count < ans {
		ans = count
	}
	for i := 1; i < l-k+1; i++ {
		if blocks[i-1] == 'W' {
			count--
		}
		if blocks[i+k-1] == 'W' {
			count++
		}
		if count < ans {
			ans = count
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	str := "WBBWWBBWBW"
	result := minimumRecolors(str, 7)
	fmt.Println(result)
}
