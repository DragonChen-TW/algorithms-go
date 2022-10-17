package main

import "fmt"

// Author:	DragonChen https://github.com/dragonchen-tw/
// Title:	Ch2 ex1-1 - Insertion Sort
// Date:	2022/05/23

// Textbook p24
func InsertionSort(nums []int, n int) {
	// i, j change
	for i := 1; i < n; i++ {
		for k := 0; k < n; k++ {
			if k == i {
				fmt.Printf(" vv")
			} else {
				fmt.Printf("   ")
			}
		}
		fmt.Println()
		fmt.Println(nums)

		key := nums[i]
		// insert key into correct position in nums[0:i-1]
		j := i - 1
		for j >= 0 && nums[j] > key {
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
