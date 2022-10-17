package main

import "fmt"

// Author:	DragonChen https://github.com/dragonchen-tw/
// Title:	Ch2 ex1-3 - Insertion Sort (Decreasing)
// Date:	2022/10/14

// Textbook p25
func InsertionSort(nums []int, n int) {
	// i, j change
	for i := 1; i < n; i++ {
		key := nums[i]
		// insert key into correct position in nums[0:i-1]
		j := i - 1
		for j >= 0 && nums[j] <= key {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
	}
}

func main() {
	var a []int = []int{31, 41, 59, 26, 41, 58}
	InsertionSort(a, 6)
	fmt.Println(a)
}
