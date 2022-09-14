package main

import "fmt"

//27.移除元素.remove-element

// leetcode submit region begin(Prohibit modification and deletion)
func removeElement(nums []int, val int) int {
	l := len(nums)
	ans := l
	p, q := 0, 0
	for q < l {
		if nums[q] != val {
			nums[p] = nums[q]
			p++
			q++
		} else if nums[q] == val {
			ans--
			q++
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	a := []int{3, 1, 1, 3, 2, 3, 3, 4, 4, 5, 6, 7, 7}
	result := removeElement(a, 3)
	fmt.Println(result)
	fmt.Println(a)
}
