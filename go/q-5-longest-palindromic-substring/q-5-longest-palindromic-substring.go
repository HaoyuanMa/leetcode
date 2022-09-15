package main

import "fmt"

//5.最长回文子串.longest-palindromic-substring

// leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	var ans string
	var cur string
	for i := 0; i < len(s); i++ {
		sin := getTmpAns(s, i, i)
		dou := getTmpAns(s, i, i+1)
		if len(sin) > len(dou) {
			cur = sin
		} else {
			cur = dou
		}
		if len(cur) > len(ans) {
			ans = cur
		}
	}
	return ans
}

func getTmpAns(s string, p, q int) string {
	l := len(s)
	for p >= 0 && q < l && s[p] == s[q] {
		p--
		q++
	}

	return s[p+1 : q]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	a := "abcdcbafg"
	result := longestPalindrome(a)
	fmt.Println(result)
}
