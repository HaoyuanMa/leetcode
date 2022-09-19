package main

import "fmt"

//509.斐波那契数.fibonacci-number

// leetcode submit region begin(Prohibit modification and deletion)
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	head, tail := 0, 1
	//两步两步走
	//for i := 0; i < n/2; i++ {
	//	head = head + tail
	//	tail = head + tail
	//}
	//if n%2 == 1 {
	//	return tail
	//} else {
	//	return head
	//}

	//一步一步走
	for i := 0; i < n-1; i++ {
		tail = head + tail
		head = tail - head
	}
	return tail
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := fib(3)
	fmt.Println(result)
}
