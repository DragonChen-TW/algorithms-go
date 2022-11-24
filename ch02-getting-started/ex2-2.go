package main

import "fmt"

// Author:	DragonChen https://github.com/dragonchen-tw/
// Title:	Part1 Ch2 ex2-1 Selection Sort
// Date:	2022/11/02

func main() {
	fmt.Println(selectionSort([]int{4, 2, 1, 6, 3, 9}))
	fmt.Println(selectionSort([]int{2, 9, 5, 6, 8, 1, 3}))
	fmt.Println(selectionSort([]int{5, 4, 3, 2, 1}))
}

/*
Inintialization: i starts from 0, nums[:i] is sorted (there is 0 elements are there right now).
Maintence: iterate through nums[i:], record the index of smallest number to variabel minIndex.
    Swap nums[i] and nums[minIndex]. i++.
Terminatence: i stops len(nums) - 1.
*/

func selectionSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			n := nums[j]
			if n < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}

	return nums
}
