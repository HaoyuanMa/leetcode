package main

import (
	"fmt"
)

//1109.航班预订统计.corporate-flight-bookings

// leetcode submit region begin(Prohibit modification and deletion)
type Diff struct {
	nums []int
	size int
}

func (d *Diff) Init(a []int, n int) {
	d.size = n
	d.nums = make([]int, n)
	if len(a) == 0 {
		return
	}
	d.nums[0] = a[0]
	for i := 1; i < d.size; i++ {
		d.nums[i] = a[i] - a[i-1]
	}
}
func (d *Diff) Update(i, j, dt int) {
	d.nums[i] += dt
	if j < d.size-1 {
		d.nums[j+1] -= dt
	}
}
func (d *Diff) Result() []int {
	res := make([]int, d.size)
	res[0] = d.nums[0]
	for i := 1; i < d.size; i++ {
		res[i] = res[i-1] + d.nums[i]
	}
	return res
}
func corpFlightBookings(bookings [][]int, n int) []int {
	d := &Diff{}
	d.Init(nil, n)
	for _, v := range bookings {
		d.Update(v[0]-1, v[1]-1, v[2])
	}
	return d.Result()
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	in := [][]int{
		{1, 2, 10},
		{2, 3, 20},
		{2, 5, 25},
	}
	result := corpFlightBookings(in, 5)
	fmt.Println(result)
}
