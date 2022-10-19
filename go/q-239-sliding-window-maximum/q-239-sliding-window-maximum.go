package main

import (
	"container/heap"
	"fmt"
)

//239.滑动窗口最大值.sliding-window-maximum

// leetcode submit region begin(Prohibit modification and deletion)
type bigHeap []int

type pq *bigHeap

func (h bigHeap) Len() int {
	return len(h)
}
func (h bigHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h bigHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h *bigHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *bigHeap) Pop() interface{} {
	out := (*h)[0]
	*h = (*h)[:len(*h)-1]
	return out
}

// const MIN = math.MinInt
func maxSlidingWindow(nums []int, k int) []int {
	a := []int{}
	h := (*bigHeap)(&a)
	for i := 0; i < k; i++ {
		*h = append(*h, nums[i])
	}
	heap.Init(h)
	sz := len(nums)
	ans := []int{}
	for i := 0; i <= sz-k; i++ {
		out := (*h)[0]
		ans = append(ans, out)
		if i == sz-k {
			break
		}
		toOut := nums[i]
		toIn := nums[i+k]
		heap.Push(h, toIn)
		for j := 0; j < len(*h); j++ {
			if (*h)[j] == toOut {
				heap.Remove(h, j)
				break
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	fmt.Println(result)
}
