package main

import "fmt"

// Author:	DragonChen https://github.com/dragonchen-tw/
// Title:	Part1 Ch2 3-6 Binary Search
// Date:	2022/11/24

/*
Procedure:
BINARY_SEARCH(nums, target) {
	l, r = 0, len(nums)-1

	while l <= r:
	    mid = (l + r) / 2
		if nums[mid] < target:
		    l = mid + 1
		else if nums[mid] > target:
		    r = mid - 1
		else:
			return mid
	return l
}
*/

// Textbook p44
func BinarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	a := []int{26, 31, 41, 59, 63, 87, 93}
	fmt.Println(BinarySearch(a, 59)) // 3
	fmt.Println(BinarySearch(a, 60)) // -1
	fmt.Println(BinarySearch(a, 87)) // 5
	// fmt.Println(a)
}
