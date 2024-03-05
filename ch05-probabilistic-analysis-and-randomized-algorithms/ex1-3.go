package main

import (
	"fmt"
	"math/rand"
)

// Author:	DragonChen https://github.com/dragonchen-tw/
// Title:	Ex 1.3 Generate uniform random numbers by a biased random
// Date:	2024/03/05

func main() {
	count, total := 0, 0
	for i := 0; i < 1000; i++ {
		if Random() == 0 {
			count++
		}
		total++
	}
	ratio := float64(count) / float64(total)
	fmt.Println("count:", count, "ratio", ratio)
}

// By running BiasedRandom() twice, we compare their results to generate a uniform random number
// The runtime of Random() is O(2 * p*2 + (1 - 2p) * 2 * (p-2p*2) * 2)
func Random() int {
	// return 0 with probability 1/2, and 1 with probability 1/2
	// by running BiasedRandom() as subroutine
	q1, q2 := BiasedRandom(), BiasedRandom()
	if q1 == 0 && q2 == 1 {
		return 0
	} else if q1 == 1 && q2 == 0 {
		return 1
	} else {
		return Random()
	}
}

func BiasedRandom() int {
	// return 0 with probability 1/3, and 1 with probability 2/3
	var p float64 = 1 / 3.0

	if rand.Float64() < p {
		return 0
	} else {
		return 1
	}
}
