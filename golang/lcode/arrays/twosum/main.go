package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 7, 11, 15}
	target := 9

	res := TwoSum(arr, target)
	fmt.Println(res)
}

// TwoSum given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// Given nums = [2, 7, 11, 15], target = 9,
// iterate through the list of numbers, find its complement, by substracting the num from the target,
// store the complement's index location (num required to add to get the target) it in a k/v store
// for each value we iterate through, check if it's complement is in the map, if so, return both indexes
func TwoSum(nums []int, target int) []int {
	complementMap := make(map[int]int)

	for i, num := range nums {
		idx, found := complementMap[num]
		if found {
			return []int{idx, i}
		}
		complement := target - num
		complementMap[complement] = i
		fmt.Println("complement map:", complementMap)
	}

	return []int{}
}
