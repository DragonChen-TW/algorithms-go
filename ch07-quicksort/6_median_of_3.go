package main

import (
	"fmt"
	"math/rand"
)

// Author:	DragonChen https://github.com/dragonchen-tw/
// Title:	Part02 ch07 Quicksort 6 median of 3
// Date:	2024/04/22

func main() {
	A := []int{2, 8, 7, 1, 3, 5, 6, 4}
	Quicksort(A, 0, len(A)-1)
	fmt.Println("A", A)
}

func Quicksort(A []int, p, r int) {
	if p < r {
		q := Partition(A, p, r)
		Quicksort(A, p, q-1)
		Quicksort(A, q+1, r)
	}
}

func Partition(A []int, p, r int) int {
	if len(A) > 10 {
		mid := (p + r) / 2
		if (A[mid] >= A[p] && A[mid] <= A[r]) || (A[mid] <= A[p] && A[mid] >= A[r]) {
			A[mid], A[r] = A[r], A[mid]
		} else if (A[p] >= A[mid] && A[p] <= A[r]) || (A[p] <= A[mid] && A[p] >= A[r]) {
			A[p], A[r] = A[r], A[p]
		}
		// else means A[r] is the median, no need to exchange
	}

	x := A[r] // use the rightmost element as pivot
	i := p - 1
	for j := p; j < r; j++ {
		if A[j] <= x {
			i++
			A[i], A[j] = A[j], A[i]
		}
	}
	A[i+1], A[r] = A[r], A[i+1]
	return i + 1
}

// 7-3
func RandomizedPartition(A []int, p, r int) int {
	idx := rand.Intn(r-p+1) + p
	A[r], A[idx] = A[idx], A[r]
	return Partition(A, p, r)
}

func RandomizedQuicksort(A []int, p, r int) {
	if p < r {
		q := RandomizedPartition(A, p, r)
		RandomizedQuicksort(A, p, q-1)
		RandomizedQuicksort(A, q+1, r)
	}
}
