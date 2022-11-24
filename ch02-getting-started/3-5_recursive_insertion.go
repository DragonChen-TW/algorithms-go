package main

import "fmt"

// Author:	DragonChen https://github.com/dragonchen-tw/
// Title:	Part1 Ch2 - 3-5 Recursive Insertion Sort
// Date:	2022/11/24

/*
Procedure:
RECURSIVE_INSERTION_SORT(nums, n) {
	if n == 0:
		return

	RECURSIVE_INSERTION_SORT(nums, n-1)

	i = n - 1
	while i > 0 && nums[i] < nums[i-1]:
		temp = nums[i]
		nums[i] = nums[i-1]
		nums[i-1] = temp
		i--
}
*/

// Textbook p44
func RecursiveInsertionSort(nums []int) {
	if len(nums) == 1 {
		return
	}

	// sort nums[0:length-2]
	var length int = len(nums)
	RecursiveInsertionSort(nums[:length-1])

	// insert nums[length-1] into sorted nums[0:length-2]
	i := length - 1
	for i > 0 && nums[i] < nums[i-1] {
		nums[i], nums[i-1] = nums[i-1], nums[i]
		i--
	}
}

func main() {
	var a []int = []int{5, 2, 4, 6, 1, 3}
	RecursiveInsertionSort(a)
	fmt.Println(a)
}
