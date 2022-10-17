package main

import (
	"fmt"
)

//3.无重复字符的最长子串.longest-substring-without-repeating-characters

// leetcode submit region begin(Prohibit modification and deletion)
const MAX = 130

type win struct {
	countAlpha [MAX]int
	countValid int
	countNums  int
}

func (w *win) init() {
	w.countValid = MAX
	w.countNums = 0
}
func (w *win) isValid() bool {
	return w.countValid == MAX
}
func (w *win) push(ch byte) {
	w.countNums++
	w.countAlpha[ch]++
	if w.countAlpha[ch] == 2 {
		w.countValid--
	}
}
func (w *win) pop(ch byte) {
	if w.countAlpha[ch] == 2 {
		w.countValid++
	}
	w.countNums--
	w.countAlpha[ch]--
}

func lengthOfLongestSubstring(s string) int {
	w := &win{}
	w.init()
	ans := 0
	for left, right := 0, 0; right < len(s); right++ {
		toIn := s[right]
		w.push(toIn)
		for ; !w.isValid() && left < right; left++ {
			toOut := s[left]
			w.pop(toOut)
		}
		if w.countNums > ans {
			ans = w.countNums
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := lengthOfLongestSubstring("asfsfsfsfdfsd")
	fmt.Println(result)
}
