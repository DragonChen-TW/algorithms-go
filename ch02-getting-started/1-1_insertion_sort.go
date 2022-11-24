package main

import "fmt"

// Author:	DragonChen https://github.com/dragonchen-tw/
// Title:	Part1 Ch2 - 1-1 Insertion Sort
// Date:	2022/05/23

// Textbook p19
func InsertionSort(nums []int, n int) {
	// i, j change
	for i := 1; i < n; i++ {
		key := nums[i]
		// insert key into correct position in nums[0:i-1]
		j := i - 1
		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
		// fmt.Println(nums)
	}
}

func main() {
	var a []int = []int{5, 2, 4, 6, 1, 3}
	InsertionSort(a, 6)
	fmt.Println(a)
}
