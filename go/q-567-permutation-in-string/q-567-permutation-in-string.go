package main

import (
	"fmt"
)

//567.字符串的排列.permutation-in-string

// leetcode submit region begin(Prohibit modification and deletion)
type counter struct {
	alpha [TOTAL]int
	zero  int
}

const (
	PUSH = iota
	POP
	TOTAL = 130
)

func (c *counter) change(ch byte, op int) {
	if c.alpha[ch] == 0 {
		c.zero--
	}
	if op == PUSH {
		c.alpha[ch]++
	} else {
		c.alpha[ch]--
	}
	if c.alpha[ch] == 0 {
		c.zero++
	}
}

// clearer logic
func checkInclusion(s1 string, s2 string) bool {
	c := &counter{zero: TOTAL}
	l1 := len(s1)
	l2 := len(s2)
	if l1 > l2 {
		return false
	}
	for i := 0; i < l1; i++ {
		c.change(s1[i], POP)
	}
	for i := 0; i < l1; i++ {
		c.change(s2[i], PUSH)
	}
	for i := 0; i+l1 < l2 && c.zero != TOTAL; i++ {
		out := s2[i]
		in := s2[i+l1]
		c.change(out, POP)
		c.change(in, PUSH)
	}
	if c.zero == TOTAL {
		return true
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := checkInclusion("adc", "dcda")
	fmt.Println(result)
}

func checkInclusion_yzh(s1, s2 string) bool {
	countA := [130]int{}
	count0 := 130
	l1 := len(s1)
	l2 := len(s2)
	for i := 0; i < l1; i++ {
		if countA[s1[i]] == 0 {
			count0--
		}
		countA[s1[i]]--
	}
	for i := 0; i < l1; i++ {
		countA[s2[i]]++
	}
	for i := 0; i+l1-1 < l2 && count0 != 130; i++ {
		out := s2[i]
		in := s2[i+l1-1]
		countA[out]--
		countA[in]++
		if countA[out] == 0 {
			count0--
		}
		if countA[in] == 0 {
			count0++
		}
	}
	if count0 == 130 {
		return true
	}
	return false
}

func checkInclusion_lbld(s1 string, s2 string) bool {
	need := map[byte]int{}
	has := map[byte]int{}
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}
	szt := len(need)
	for l, r, count := 0, 0, 0; r < len(s2); {

		if need[s2[r]] > 0 {
			has[s2[r]]++
			if has[s2[r]] == need[s2[r]] {
				count++
			}
		}
		for r-l+1 > len(s1) {
			if need[s2[l]] > 0 {
				if has[s2[l]] == need[s2[l]] {
					count--
				}
				has[s2[l]]--
			}
			l++
		}
		if count == szt {
			return true
		}
		r++
	}
	return false
}
