package main

import "fmt"

// Author:	DragonChen https://github.com/dragonchen-tw/
// Title:	Part1 Ch2 - 3-1 Insertion Sort
// Date:	2022/11/02

// Textbook p44
func MergeSort(nums []int, p, r int) {
	fmt.Println(p, r)
	if p >= r {
		return
	}
	q := (p + r) / 2
	MergeSort(nums, p, q)
	MergeSort(nums, q+1, r)
	Merge(nums, p, q, r)
}
func Merge(nums []int, p, q, r int) {
	lenL := q - p + 1
	lenR := r - q
	var L []int = make([]int, lenL)
	var R []int = make([]int, lenR)
	copy(L, nums[p:q+1])
	copy(R, nums[q+1:r+1])

	var i, j, k int = 0, 0, p
	for i < lenL && j < lenR {
		if L[i] <= R[j] {
			nums[k] = L[i]
			i++
		} else {
			nums[k] = R[j]
			j++
		}
		k++
	}

	for i < lenL {
		nums[k] = L[i]
		i++
		k++
	}
	for j < lenR {
		nums[k] = R[j]
		j++
		k++
	}
}

func main() {
	var a []int
	// a = []int{5, 2, 4, 6, 1, 3}
	// MergeSort(a, 0, len(a)-1)
	// fmt.Println(a)
	a = []int{12, 3, 7, 9, 14, 6, 11, 2}
	MergeSort(a, 0, len(a)-1)
	fmt.Println(a)
}
