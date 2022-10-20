package main

import (
	"container/heap"
	"fmt"
)

//239.滑动窗口最大值.sliding-window-maximum

// leetcode submit region begin(Prohibit modification and deletion)
type item struct {
	Val   int
	Index int
	Count int
}

type bigHeap []*item

func (h bigHeap) Len() int {
	return len(h)
}
func (h bigHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].Index = i
	h[j].Index = j
}
func (h bigHeap) Less(i, j int) bool {
	return h[i].Val > h[j].Val
}
func (h *bigHeap) Push(x interface{}) {
	sz := len(*h)
	it := x.(*item)
	it.Index = sz
	*h = append(*h, it)
}
func (h *bigHeap) Pop() interface{} {
	old := *h
	n := len(old)
	out := old[n-1]
	old[n-1] = nil
	out.Index = -1
	*h = old[:n-1]
	return out
}

// const MIN = math.MinInt
func maxSlidingWindow(nums []int, k int) []int {
	a := []*item{}
	h := (*bigHeap)(&a)
	tr := map[int]*item{}
	for i := 0; i < k; i++ {
		if v, ok := tr[nums[i]]; ok {
			v.Count++
		} else {
			toAdd := &item{nums[i], len(*h), 1}
			tr[nums[i]] = toAdd
			*h = append(*h, toAdd)
		}
	}
	heap.Init(h)
	sz := len(nums)
	ans := []int{}
	for i := 0; i <= sz-k; i++ {
		out := (*h)[0].Val
		ans = append(ans, out)
		if i == sz-k {
			break
		}
		toOut := nums[i]
		toIn := nums[i+k]
		if v, ok := tr[toIn]; ok {
			v.Count++
		} else {
			toAdd := &item{toIn, len(*h), 1}
			heap.Push(h, toAdd)
			tr[toIn] = toAdd
		}
		if v, ok := tr[toOut]; ok {
			v.Count--
			if v.Count == 0 {
				delete(tr, v.Val)
				heap.Remove(h, v.Index)
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := maxSlidingWindow([]int{1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 3)
	fmt.Println(result)
}
