package main

import "fmt"

//167.两数之和 II - 输入有序数组.two-sum-ii-input-array-is-sorted

// leetcode submit region begin(Prohibit modification and deletion)
func twoSum(numbers []int, target int) []int {
	p, q := 0, len(numbers)-1
	for p != q {
		tmp := numbers[p] + numbers[q]
		if tmp == target {
			return []int{p + 1, q + 1}
		} else if tmp < target {
			p++
		} else {
			q--
		}
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := twoSum(a, 15)
	fmt.Println(result)
}
