package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("List lenght:", len(list))
	result := binarySearch(list, 8)
	fmt.Printf("result: %v\n", result)
}

func binarySearch(list []int, target int) int {
	left, right := 0, len(list)-1
	for left <= right {
		i := (left + right) / 2
		switch {
		case list[i] == target:
			return i
		case list[i] > target:
			right = i - 1
		case list[i] < target:
			left = i + 1
		}
	}
	return -1
}
