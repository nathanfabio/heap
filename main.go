package main

import "fmt"

func main() {
	m := &MaxHeap{}
	fmt.Println(m)

	buildHeap := []int{15, 30, 45, 12 ,5 ,64, 19, 21, 32}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}

type MaxHeap struct {
	arr []int
}

func (h *MaxHeap) Insert(key int) {
	h.arr = append(h.arr, key)
	h.maxHeapifyUp(len(h.arr) - 1)
}

func (h *MaxHeap) Extract() int {
	extracted := h.arr[0]
	l := len(h.arr) - 1

	if len(h.arr) == 0 {
		fmt.Println("Heap is empty")
		return -1
	}

	h.arr[0] = h.arr[len(h.arr) - 1]
	h.arr = h.arr[:l]

	h.maxHeapifyDown(0)

	return extracted
}
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.arr[parent(index)] < h.arr[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.arr) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.arr[l] > h.arr[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.arr[index] < h.arr[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.arr[i1], h.arr[i2] = h.arr[i2], h.arr[i1]
}