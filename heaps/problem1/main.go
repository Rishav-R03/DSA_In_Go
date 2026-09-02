package main

/*
parent = (i - 1) / 2
left   = 2*i + 1
right  = 2*i + 2
*/

// heap is a complete binary tree stored in an array.

type MinHeap struct {
	data []int
}

func (h *MinHeap) Insert(val int) {
	h.data = append(h.data, val)
	h.heapifyUp(len(h.data) - 1)
}

func (h *MinHeap) heapifyUp(ind int) {
	for ind > 0 {
		parent := (ind - 1) / 2

		if h.data[parent] <= h.data[ind] {
			break
		}
		h.data[parent], h.data[ind] = h.data[ind], h.data[parent]
		ind = parent
	}
}

func (h *MinHeap) ExtractMin() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}
	min := h.data[0]
	last := len(h.data) - 1

	h.data[0] = h.data[last]
	if len(h.data) > 0 {
		h.heapifyUp(0)
	}
	return min, true
}
