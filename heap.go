package main

import (
    "fmt"
)

type MinHeap struct {
    arr []int
}

func NewMinHeap() *MinHeap {
    return &MinHeap {
        arr: make([]int, 0),
    }
}

func (h *MinHeap) Insert(value int) {
    h.arr = append(h.arr, value)
    h.heapifyUp(len(h.arr) - 1)
}

func (h *MinHeap) ExtractMin() int {
    if h.IsEmpty() {
        panic("Heap is empty")
    }
    min := h.arr[0]
    lastIndex := len(h.arr) - 1
    h.arr[0] = h.arr[lastIndex]
    h.arr = h.arr[:lastIndex]
    h.heapifyDown(0)
    return min
}

func (h *MinHeap) IsEmpty() bool {
    return len(h.arr) == 0
}

func (h *MinHeap) heapifyUp(index int) {
    parentIndex := (index - 1) / 2
    for index > 0 && h.arr[index] < h.arr[parentIndex] {
		h.arr[index], h.arr[parentIndex] = h.arr[parentIndex], h.arr[index]
        index = parentIndex
        parentIndex = (index - 1) / 2
    }
}

func (h *MinHeap) heapifyDown(index int) {
    for {
        leftChildIndex := 2*index + 1
        rightChildIndex := 2*index + 2
        smallestIndex := index

		if leftChildIndex < len(h.arr) && h.arr[leftChildIndex] < h.arr[smallestIndex] {    
            smallestIndex = leftChildIndex
        }

		if rightChildIndex < len(h.arr) && h.arr[rightChildIndex] < h.arr[smallestIndex] {
            smallestIndex = rightChildIndex
        }

        if smallestIndex == index {
            break
        }

		h.arr[index], h.arr[smallestIndex] = h.arr[smallestIndex], h.arr[index]
        index = smallestIndex
    }
}

func main() {
    heap := NewMinHeap()
    
    heap.Insert(10)
    heap.Insert(5)
    heap.Insert(7)
    heap.Insert(3)
    heap.Insert(12)

    for !heap.IsEmpty() {
        min := heap.ExtractMin()
        fmt.Println(min)
    }
}
