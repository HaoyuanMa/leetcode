package main

import "fmt"

//12.整数转罗马数字.integer-to-roman

// leetcode submit region begin(Prohibit modification and deletion)
func intToRoman(num int) string {
	base := [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman := [13]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	ans := ""
	for i := 0; i < 13; i++ {
		count := num / base[i]
		for j := 0; j < count; j++ {
			ans += roman[i]
		}
		num -= count * base[i]
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := intToRoman(99)
	fmt.Println(result)
}
