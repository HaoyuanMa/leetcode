package main

import "fmt"

//26.删除有序数组中的重复项.remove-duplicates-from-sorted-array

// leetcode submit region begin(Prohibit modification and deletion)
func removeDuplicates(nums []int) int {
	l := len(nums)
	ans := l
	p, q := 0, 1
	for q < l {
		if nums[p] == nums[q] {
			ans--
		} else {
			p++
			nums[p] = nums[q]
		}
		q++
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	a := []int{1, 1, 2, 3, 3, 4, 4, 5, 6, 7, 7}
	result := removeDuplicates(a)
	fmt.Println(result)
	fmt.Println(a)
}
