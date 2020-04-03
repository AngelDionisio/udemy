package main

import (
	"fmt"
	"sort"
)

type test struct {
	list   []int
	target int
}

func main() {
	input := []int{10, 2, 9, 12, 1}
	target := 10
	findPair(input, target)

	// tests := []test{
	// 	{
	// 		list:   []int{1, 2, 5},
	// 		target: 10,
	// 	},
	// 	{
	// 		list:   []int{1, 2, 9, 12},
	// 		target: 10,
	// 	},
	// 	{
	// 		list:   []int{1, 9, 1, 9},
	// 		target: 10,
	// 	},
	// 	{
	// 		list:   []int{1, 9, 1},
	// 		target: 10,
	// 	},
	// 	{
	// 		list:   []int{1, 2, 9, 12, 10},
	// 		target: 10,
	// 	},
	// }

	// for _, v := range tests {
	// 	findPair(v.list, v.target)
	// }
}

// printPairsWhoseSumIsN brute force
func printPairsWhoseSumIsN(list []int, target int) {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i]+list[j] == target {
				fmt.Printf("The sum of %d + %d = %d\n", list[i], list[j], target)
			}
		}
	}
}

// more efficient way to find sum of pairs
/*
* check if list is sorted, if not, sort
* create two indexes, lo, hi, pointing to the start and tail of the list
* check if the value at arr[hi] > target, if so decrease the value of tail (hi--)
* check arr[lo] + arr[hi] == sum, if true, add arr[lo] and arr[hi] to results array
* else if sum < target, lo++
* else if sum > target, hi--
* remove any duplicates from results array
 */
func findPair(list []int, target int) {
	// check if array is sorted, if not, sort ascending
	if !sort.IntsAreSorted(list) {
		sort.Ints(list)
	}

	// keep two indices pointing to the end-points of the array
	lo := 0
	hi := len(list) - 1
	var l []int

	for lo < hi {
		if list[hi] > target {
			hi--
			continue
		}

		tempSum := list[lo] + list[hi]
		if tempSum == target {
			// fmt.Println("target sum found:", sortedInts[lo], " + ", sortedInts[hi], "= ", target)
			l = append(l, list[lo])
			l = append(l, list[hi])
			lo++
			hi--
			continue
		} else if tempSum < target {
			lo++
		} else if tempSum > target {
			hi--
		}
	}

	result := removeDuplicates(l)

	if len(result) == 0 {
		fmt.Println("No pairs found")
	}

	fmt.Println(result)
}

func removeDuplicates(list []int) []int {
	seenMap := make(map[int]bool)
	uniqueItems := []int{}

	for _, v := range list {
		if _, found := seenMap[v]; !found {
			seenMap[v] = true
			uniqueItems = append(uniqueItems, v)
		}
	}

	return uniqueItems
}
