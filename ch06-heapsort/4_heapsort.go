package main

import "fmt"

// Author:	DragonChen https://github.com/dragonchen-tw/
// Title:	Part02 ch06 Heapsort
// Date:	2024/03/06

func main() {
	data := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	A := NewHeap(data)
	BuildMaxHeap(A, A.size)
	// fmt.Println(A)

	Heapsort(A, A.size)
	fmt.Println(A.data[1:])
}

type Heap struct {
	data []int
	size int
}

func NewHeap(data []int) Heap {
	data = append([]int{-1}, data...)
	return Heap{
		data: data,
		size: len(data) - 1,
	}
}

// Figure 6.2
func MaxHeapify(A Heap, i int) {
	l := i << 1 // left child of i
	r := l + 1  // right child of i

	var largest int
	if l <= A.size && A.data[l] > A.data[i] {
		largest = l
	} else {
		largest = i
	}
	if r <= A.size && A.data[r] > A.data[largest] {
		largest = r
	}

	if largest != i {
		A.data[i], A.data[largest] = A.data[largest], A.data[i]
		MaxHeapify(A, largest)
	}
}

// 6.3
func BuildMaxHeap(A Heap, n int) {
	A.size = n
	for i := n >> 1; i >= 1; i-- {
		MaxHeapify(A, i)
	}
}

// 6.4
func Heapsort(A Heap, n int) {
	BuildMaxHeap(A, n)
	for i := n; i >= 2; i-- {
		A.data[1], A.data[i] = A.data[i], A.data[1]
		A.size--
		MaxHeapify(A, 1)
	}
}

// 6.5
func MaxHeapMaximum(A Heap) (int, error) {
	if A.size < 1 {
		return 0, fmt.Errorf("heap underflow")
	}
	return A.data[1], nil
}

func MaxHeapExtractMax(A Heap) int {
	max := A.data[1]
	A[1] = A.data[A.size]
	A.size--
	MaxHeapify(A, 1)
	return max
}

// 6.5
func MaxHeapIncreaseKey(A Heap, i, key int) error {
	if key < A.data[i] {
		return fmt.Errorf("new key is smaller than current key")
	}
	A.data[i] = key
	for i > 1 && a.data[i>>1] < A.data[i] {
		A.data[i], A.data[i>>1] = A.data[i>>1], A.data[i]
		i = i >> 1
	}
	return nil
}

func MaxHeapInsert(A Heap, n int) error {
	A.size++
	MaxHeapIncreaseKey(A, A.size, n)
}
