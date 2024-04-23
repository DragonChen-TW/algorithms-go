package main

import "fmt"

// Author:	DragonChen https://github.com/dragonchen-tw/
// Title:	Part02 ch07 Stooge sort
// Date:	2024/03/07

func main() {
	A := []int{2, 8, 7, 1, 3, 5, 6, 4}
	Quicksort(A, 0, len(A)-1)
	fmt.Println("A", A)
}

func StoogeSort(A []int, p, r int) {
	if A[p] > A[r] {
		A[p], A[r] = A[r], A[p]
	}
	if p+1 < r {
		k = (r - p + 1) / 3
		StoogeSort(A, p, r-k)
		StoogeSort(A, p+k, r)
		StoogeSort(A, p, r-k)
	}
}
