package main

import "fmt"

// Author:	DragonChen https://github.com/dragonchen-tw/
// Title:	Part1 Ch2 ex1-4 Linear Search
// Date:	2022/10/14

// Textbook version p25
func LinearSearch(nums []int, value int) int {
	i := 0
	for i < len(nums) && nums[i] != value {
		i++
	}

	if i >= len(nums) {
		return -1
	}
	return i
}

func main() {
	a := []int{31, 41, 59, 26, 41, 58}
	fmt.Println(LinearSearch(a, 58)) // 5
	fmt.Println(LinearSearch(a, 60)) // -1
	fmt.Println(LinearSearch(a, 31)) // 0
	// fmt.Println(a)
}
