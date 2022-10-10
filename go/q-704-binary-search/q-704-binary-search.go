package main

import "fmt"

//704.二分查找.binary-search

// leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := search([]int{1, 2, 3, 4, 5}, 3)
	fmt.Println(result)
}
