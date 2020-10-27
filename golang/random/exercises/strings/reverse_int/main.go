package main

import (
	"fmt"
	"math"
)

/*
Given a 32-bit signed integer, reverse digits of an integer.

Note:
Assume we are dealing with an environment that could only store
integers within the 32-bit signed integer range: [−231,  231 − 1].
For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.

Input: x = 123
Output: 321
Example 2:

Input: x = -123
Output: -321
Example 3:

Input: x = 120
Output: 21
Example 4:

Input: x = 0
Output: 0

*/

func main() {
	nums := []int{123, -123, 120, 0}
	for _, num := range nums {
		fmt.Printf("reverse(%v): %v\n", num, reverse(num))
	}
}

func reverse(x int) int {
	reversed := 0

	for pop := 0; x != 0; x /= 10 {
		pop = x % 10                     // pluck last int from number
		reversed = (reversed * 10) + pop // add an extra zero to reversed, add plucked digit to end
	}

	if reversed < math.MinInt32 || reversed > math.MaxInt32 {
		return 0
	}

	return reversed
}
