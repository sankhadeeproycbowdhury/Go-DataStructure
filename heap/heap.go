package heap

import (
	"container/heap"
	"fmt"
)

type HeapItem[T any] struct {
	Priority int
	Data     T
}

type GenericHeap[T any] []HeapItem[T]

// Implementation of heap.Interface/sort.Interface
func (h GenericHeap[T]) Len() int           { return len(h) }
func (h GenericHeap[T]) Less(i, j int) bool { return h[i].Priority < h[j].Priority }
func (h GenericHeap[T]) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *GenericHeap[T]) Push(x any) {
	*h = append(*h, x.(HeapItem[T]))
}

func (h *GenericHeap[T]) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

// --- NEW METHODS ---

func NewHeap[T any](items ...HeapItem[T]) *GenericHeap[T] {
	h := GenericHeap[T](items)
	heap.Init(&h)
	return &h
}

func (h *GenericHeap[T]) Peek() (HeapItem[T], bool) {
	if h.Len() == 0 {
		return HeapItem[T]{}, false
	}
	return (*h)[0], true
}

func topKFrequent(nums []int, k int) []int {
    req := k
    seen := make(map[int]int)
    ans := make([]int, 0, k)

    for _, n := range nums {
        val := seen[n]
        seen[n] = val + 1
    }

	fmt.Println(seen)

    pq := NewHeap[int]()
    for k, v := range seen{
		fmt.Println(v, k)
        heap.Push(pq, HeapItem[int]{Priority: v, Data: k})
    }

    for pq.Len() > 0 && req > 0 {
        req = req - 1
        item := heap.Pop(pq).(HeapItem[int])
        ans = append(ans, item.Data)
    }

    return ans
}


