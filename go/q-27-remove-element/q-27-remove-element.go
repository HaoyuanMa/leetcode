package main

import "fmt"

//27.移除元素.remove-element

// leetcode submit region begin(Prohibit modification and deletion)
func removeElement(nums []int, val int) int {
	l := len(nums)
	ans := l
	p, q := 0, 1
	for q < l {
		if nums[p] != val {
			p++
			q++
		} else if nums[q] == val {
			ans--
			q++
		} else {
			nums[p] = nums[q]
			p++
			q++
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	a := []int{1, 1, 2, 3, 3, 4, 4, 5, 6, 7, 7}
	result := removeElement(a, 3)
	fmt.Println(result)
	fmt.Println(a)
}
