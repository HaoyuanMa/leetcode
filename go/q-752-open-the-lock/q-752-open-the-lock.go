package main

import (
	"fmt"
)

//752.打开转盘锁.open-the-lock

// leetcode submit region begin(Prohibit modification and deletion)

type pw [4]byte

func openLock(deadends []string, target string) int {
	if !notIn(deadends, "0000") {
		return -1
	}
	used := map[string]bool{}
	ans := 0
	queue := []pw{pw{'0', '0', '0', '0'}}
	used["0000"] = true
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			if size == 1 {
				queue = nil
			} else {
				queue = queue[1:]
			}
			if string(cur[:]) == target {
				return ans
			}
			for j := 0; j < 4; j++ {
				tmp := cur
				tmp[j] += 1
				if tmp[j] > '9' {
					tmp[j] = '0'
				}
				str := string(tmp[:])
				if !used[str] && notIn(deadends, str) {
					queue = append(queue, tmp)
					used[str] = true
				}
				tmp = cur

				tmp[j] -= 1
				if tmp[j] < '0' {
					tmp[j] = '9'
				}
				str = string(tmp[:])
				if !used[str] && notIn(deadends, str) {
					queue = append(queue, tmp)
					used[str] = true
				}
			}
		}
		ans++
	}
	return -1
}
func notIn(strs []string, target string) bool {
	sz := len(strs)
	for i := 0; i < sz; i++ {
		if strs[i] == target {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	target := "0202"
	result := openLock(deadends, target)
	fmt.Println(result)
}
