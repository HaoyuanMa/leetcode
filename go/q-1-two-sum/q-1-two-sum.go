package main

import (
	"fmt"
	"sort"
)

//1.两数之和.two-sum

// leetcode submit region begin(Prohibit modification and deletion)
type pair struct {
	Val   int
	Index int
}
type pairs []pair

func (p pairs) Len() int {
	return len(p)
}
func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p pairs) Less(i, j int) bool {
	return p[i].Val < p[j].Val
}

func twoSum(nums []int, target int) []int {
	sz := len(nums)
	var ps pairs
	for i := 0; i < sz; i++ {
		ps = append(ps, pair{nums[i], i})
	}
	sort.Sort(ps)
	i, j := 0, sz-1
	for i < j {
		sum := ps[i].Val + ps[j].Val
		if sum == target {
			return []int{ps[i].Index, ps[j].Index}
		} else if sum < target {
			i++
		} else {
			j--
		}
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := twoSum([]int{}, 4)
	fmt.Println(result)
}
