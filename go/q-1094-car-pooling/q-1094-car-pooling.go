package main

import "fmt"

//1094.拼车.car-pooling

// leetcode submit region begin(Prohibit modification and deletion)
type Diff struct {
	dt   []int
	size int
}

func (d *Diff) Init(a []int, n int) {
	d.size = n
	d.dt = make([]int, n)
	if len(a) == 0 {
		return
	}
	d.dt[0] = a[0]
	for i := 1; i < n; i++ {
		d.dt[i] = a[i] - a[i-1]
	}
}
func (d *Diff) Update(i, j, delta int) {
	d.dt[i] += delta
	if j < d.size-1 {
		d.dt[j+1] -= delta
	}
}
func (d *Diff) Result() []int {
	res := make([]int, d.size)
	res[0] = d.dt[0]
	for i := 1; i < d.size; i++ {
		res[i] = res[i-1] + d.dt[i]
	}
	return res
}

func carPooling(trips [][]int, capacity int) bool {
	d := &Diff{}
	d.Init(nil, 1001)
	for _, trip := range trips {
		d.Update(trip[1], trip[2]-1, trip[0])
	}
	res := d.Result()
	for _, v := range res {
		if v > capacity {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	result := carPooling(nil, 0)
	fmt.Println(result)
}
