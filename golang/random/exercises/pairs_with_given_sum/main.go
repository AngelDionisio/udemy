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
	input := []int{1, 2, 9, 12, 10}
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
// start by sorting list, compare first, and last items
// if sum is > than target, move last index to the lef
// if sum is < target, move first index to the right
func findPair(list []int, target int) {
	// check if array is sorted, if not, sort ascending
	sortedInts := list
	if !sort.IntsAreSorted(list) {
		sort.Ints(sortedInts)
	}

	// keep two indices pointing to the end-points of the array
	lo := 0
	hi := len(list) - 1
	var l []int

	for lo < hi {
		if sortedInts[hi] > target {
			hi--
			continue
		}

		tempSum := sortedInts[lo] + sortedInts[hi]
		if tempSum == target {
			// fmt.Println("target sum found:", sortedInts[lo], " + ", sortedInts[hi], "= ", target)
			l = append(l, sortedInts[lo])
			l = append(l, sortedInts[hi])
			lo++
			hi--
			continue
		} else if tempSum < target {
			lo++
		} else if tempSum > target {
			hi--
		}
	}

	result := unique(l)

	if len(result) == 0 {
		fmt.Println("No pairs found")
	}

	fmt.Println(result)
}

func unique(list []int) []int {
	record := make(map[int]bool)
	uniqueItems := []int{}

	for _, v := range list {
		if _, found := record[v]; !found {
			record[v] = true
			uniqueItems = append(uniqueItems, v)
		}
	}

	return uniqueItems
}
