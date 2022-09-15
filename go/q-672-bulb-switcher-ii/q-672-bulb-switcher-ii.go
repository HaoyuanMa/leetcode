package main

import "fmt"

//672.灯泡开关 Ⅱ.bulb-switcher-ii

// leetcode submit region begin(Prohibit modification and deletion)
func flipLights(n int, presses int) int {
	if presses == 0 {
		return 1
	}
	if n == 1 {
		return 2
	}
	if n == 2 {
		if presses == 1 {
			return 3
		} else {
			return 4
		}
	}
	if presses == 1 {
		return 4
	}
	if presses == 2 {
		return 7
	}
	return 8
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := flipLights(1, 1)
	fmt.Println(result)
}
