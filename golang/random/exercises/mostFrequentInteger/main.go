package main

import (
	"fmt"
)

/*
Given an array of integers, find the most frequent integer.
Write a method that takes an array of integers and  returns an integer
{1, 2, 3, 3} => 3
{1, 2, 3} => 1 (or 2 or 3)
{3, 1, 4, 57, 4} => 4
*/

func main() {
	nums := []int{3, 1, 4, 57, 4}
	result := mostFrequentInteger(nums)

	fmt.Println("most frequent number:", result)
}

func mostOcurrencesInMap(m map[int]int) int {
	var currentHighest int
	for k, v := range m {
		if v > currentHighest {
			currentHighest = k
		}
	}
	return currentHighest
}

func mostFrequentInteger(list []int) int {
	// initialize map with space for 15 items before reallocation
	m := make(map[int]int)

	for _, v := range list {
		// check if value exists in map, if it does not, set the seen total to 1, otherwise increment
		countTotal, found := m[v]
		if !found {
			m[v] = 1
			continue
		}
		m[v] = countTotal + 1
	}

	result := mostOcurrencesInMap(m)

	return result
}
