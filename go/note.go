package note

import (
	"fmt"
	"mhy/utils"
	"unsafe"
)

type s1 struct{ Val int }
type s2 struct{ Val int }

type ListNode = utils.SListNode

type ListNodeHeap []*ListNode

func (h ListNodeHeap) Len() int {
	return len(h)
}
func (h ListNodeHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}
func (h ListNodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ListNodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}
func (h *ListNodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func note() {
	var a interface{}
	v := a.(int)
	fmt.Println(v)

	var p1 *s1
	var p2 *s2 = (*s2)(p1)
	fmt.Println(p2)

	var d *float32
	//compile fail
	//var x *int = (*int)(d)
	var x *int = (*int)(unsafe.Pointer(d))
	fmt.Println(x)
}
