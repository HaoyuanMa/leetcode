package main

import (
	"fmt"
)

//438.找到字符串中所有字母异位词.find-all-anagrams-in-a-string

// leetcode submit region begin(Prohibit modification and deletion)
const MAX = 130

type win struct {
	countAlpha [MAX]int
	countZero  int
}

func (w *win) init(str string) {
	w.countZero = MAX
	for i := 0; i < len(str); i++ {
		ch := str[i]
		w.pop(ch)
	}
}
func (w *win) matched() bool {
	return w.countZero == MAX
}
func (w *win) push(ch byte) {
	if w.countAlpha[ch] == 0 {
		w.countZero--
	}
	w.countAlpha[ch]++
	if w.countAlpha[ch] == 0 {
		w.countZero++
	}
}
func (w *win) pop(ch byte) {
	if w.countAlpha[ch] == 0 {
		w.countZero--
	}
	w.countAlpha[ch]--
	if w.countAlpha[ch] == 0 {
		w.countZero++
	}
}
func findAnagrams(s string, p string) []int {
	ans := []int{}
	if len(s) < len(p) {
		return ans
	}
	w := &win{}
	w.init(p)
	szp := len(p)
	for i := 0; i < szp; i++ {
		w.push(s[i])
	}
	for i := 0; i+szp-1 < len(s); i++ {
		if w.matched() {
			ans = append(ans, i)
		}
		toPop := s[i]
		w.pop(toPop)
		if i+szp < len(s) {
			toPush := s[i+szp]
			w.push(toPush)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := findAnagrams("cbaebabacd", "abc")
	fmt.Println(result)
}
