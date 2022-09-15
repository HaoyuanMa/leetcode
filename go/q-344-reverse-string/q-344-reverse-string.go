package main

import "fmt"

//344.反转字符串.reverse-string

// leetcode submit region begin(Prohibit modification and deletion)
func reverseString(s []byte) {
	p, q := 0, len(s)-1
	for p < q {
		tmp := s[p]
		s[p] = s[q]
		s[q] = tmp
		p++
		q--
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	a := []byte("abcdefg")
	reverseString(a)
	fmt.Println(a)
}
