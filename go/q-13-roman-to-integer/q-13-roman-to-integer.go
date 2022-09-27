package main

import "fmt"

//13.罗马数字转整数.roman-to-integer

// leetcode submit region begin(Prohibit modification and deletion)
func romanToInt(s string) int {
	base := map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}
	l := len(s)
	ans := 0
	for i := 0; i < l; i++ {
		if i+1 < l && base[s[i:i+1]] < base[s[i+1:i+2]] {
			ans += base[s[i:i+2]]
			i++
		} else {
			ans += base[s[i:i+1]]
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := romanToInt("IX")
	fmt.Println(result)
}
