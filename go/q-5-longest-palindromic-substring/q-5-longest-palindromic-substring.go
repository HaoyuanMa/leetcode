package main

import "fmt"

//5.最长回文子串.longest-palindromic-substring

// leetcode submit region begin(Prohibit modification and deletion)

var memo [1001][1001]bool

func longestPalindrome(s string) string {
	l := len(s)
	var ans string
	/*
		状态是某子串是否回文，而非某串的最大回文子串
		转移方程见 func isPalindrome()

		规划
		以长度为6的串为例，ij表示子串s[i:j+1]
		i，i和i，i+1可以直接判断，其它情况利用已得到的状态计算
		00 01 02 03 04 05
		     /  /  /  /
		   11 12 13 14 15
	            /  /  /
		      22 23 24 25
		           /  /
		         33 34 35
		              /
		            44 45
		               55
		按照从下往上，从左往右的顺序遍历所有子串并判断是否为回文
		初始状态为 55
		从此开始向上转移
	*/

	for i := l - 1; i >= 0; i-- {
		for k := i; k < l; k++ {
			if isPalindrome(s, i, k) && k-i+1 > len(ans) {
				ans = s[i : k+1]
			}
		}
	}
	return ans
}

func isPalindrome(s string, p, q int) bool {
	if p > q {
		return false
	}
	if p == q {
		memo[p][q] = true
	} else if p+1 == q && s[p] == s[q] {
		memo[p][q] = true
	} else if memo[p+1][q-1] && s[p] == s[q] {
		memo[p][q] = true
	} else {
		memo[p][q] = false
	}
	return memo[p][q]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	a := "10204"
	result := longestPalindrome(a)
	fmt.Println(result)
}

func longestPalindrome_doublePtr(s string) string {
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
