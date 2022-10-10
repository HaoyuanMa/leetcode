package main

import (
	"fmt"
)

//34.在排序数组中查找元素的第一个和最后一个位置.find-first-and-last-position-of-element-in-sorted-array

// leetcode submit region begin(Prohibit modification and deletion)
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	return []int{searchLeft(nums, target), searchRight(nums, target)}
}
func searchLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if right+1 < len(nums) && nums[right+1] == target {
		return right + 1
	} else {
		return -1
	}
}
func searchRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left-1 >= 0 && nums[left-1] == target {
		return left - 1
	} else {
		return -1
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := searchRange([]int{1}, 0)
	fmt.Println(result)
}
