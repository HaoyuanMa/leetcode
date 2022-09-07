package main

//import (
//	"container/heap"
//	"mhy/utils"
//)
//
////23.合并K个升序链表.merge-k-sorted-lists
//
//type ListNode = utils.ListNode
//
////leetcode submit region begin(Prohibit modification and deletion)
///**
//* Definition for singly-linked list.
//* type ListNode struct {
//*     Val int
//*     Next *ListNode
//* }
// */
//
//type ListNodeHeap []*ListNode
//
//func (h ListNodeHeap) Len() int {
//	return len(h)
//}
//func (h ListNodeHeap) Less(i, j int) bool {
//	return h[i].Val < h[j].Val
//}
//func (h ListNodeHeap) Swap(i, j int) {
//	h[i], h[j] = h[j], h[i]
//}
//
//func (h *ListNodeHeap) Push(x interface{}) {
//	*h = append(*h, x.(*ListNode))
//}
//func (h *ListNodeHeap) Pop() interface{} {
//	old := *h
//	n := len(old)
//	x := old[n-1]
//	*h = old[:n-1]
//	return x
//}
//
//func mergeKLists(lists []*ListNode) *ListNode {
//	n := len(lists)
//	if n == 0 {
//		return nil
//	}
//	ptrHeap := &ListNodeHeap{}
//	heap.Init(ptrHeap)
//	for i := 0; i < n; i++ {
//		if lists[i] != nil {
//			heap.Push(ptrHeap, lists[i])
//		}
//	}
//	head := &ListNode{}
//	cur := head
//	for ptrHeap.Len() > 0 {
//		t := heap.Pop(ptrHeap).(*ListNode)
//		cur.Next = t
//		cur = cur.Next
//		if t.Next != nil {
//			heap.Push(ptrHeap, t.Next)
//		}
//	}
//	return head.Next
//}
//
////leetcode submit region end(Prohibit modification and deletion)
//
//func main() {
//
//	a := utils.CreatSingleList([]int{1, 4, 5})
//	b := utils.CreatSingleList([]int{1, 3, 4})
//	c := utils.CreatSingleList([]int{2, 6})
//
//	result := mergeKLists([]*ListNode{a, b, c})
//	utils.PrintList(result)
//}
