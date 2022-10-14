package main

import (
	"fmt"
)

//76.最小覆盖子串.minimum-window-substring

// leetcode submit region begin(Prohibit modification and deletion)
const MAX = 999999

type winCounter struct {
	needCount map[byte]int
	hasCount  map[byte]int
	valid     int
}

func (wc *winCounter) need(ch byte) bool {
	return wc.needCount[ch] > 0
}
func (wc *winCounter) matched() bool {
	return wc.valid == len(wc.needCount)
}
func (wc *winCounter) push(havaIn byte) {
	if wc.need(havaIn) {
		wc.hasCount[havaIn]++
		if wc.hasCount[havaIn] == wc.needCount[havaIn] {
			wc.valid++
		}
	}
}
func (wc *winCounter) pop(toOUt byte) {
	if wc.need(toOUt) {
		if wc.hasCount[toOUt] == wc.needCount[toOUt] {
			wc.valid--
		}
		wc.hasCount[toOUt]--
	}
}
func minWindow(s string, t string) string {
	left, right := 0, 0
	szs, szt := len(s), len(t)
	ansStr, ansCt := "", MAX

	wc := &winCounter{map[byte]int{}, map[byte]int{}, 0}
	for i := 0; i < szt; i++ {
		wc.needCount[t[i]]++
	}
	for ; right < szs; right++ {
		haveIn := s[right]

		wc.push(haveIn)

		for ; wc.matched(); left++ {
			if right-left+1 < ansCt {
				ansStr = s[left : right+1]
				ansCt = len(ansStr)
			}
			toOut := s[left]
			wc.pop(toOut)
		}

	}
	if ansCt < MAX {
		return ansStr
	} else {
		return ""
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := minWindow("aa", "aa")
	fmt.Println(result)
}
