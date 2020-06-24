package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// nums := []int{1, 1, 2}

	result1 := removeDuplicates(nums)
	fmt.Printf("result: %v\n", result1)
}

func removeDuplicates(list []int) int {
	if len(list) == 0 || len(list) == 1 {
		return len(list)
	}

	left := 0
	for right := 1; right < len(list); right++ {
		fmt.Printf("left: %v, right: %v, value at left[%v] = %v\n", list[left], list[right], left, list[left])
		// since the array is sorted, if the neighbors are not equal, then
		// the item at position j is unique. At this point, we can move the
		// index of the last unique value (i), and set it to this unique item
		// which is index j
		if list[left] != list[right] {
			left++
			list[left] = list[right]
		}
	}

	fmt.Printf("after loop before ++ left: %v\n", left)
	left++
	fmt.Printf("after loop left++: %v\n", left)

	fmt.Printf("nums: %+v\n", list)
	return left
}

// func removeDuplicates(nums []int) int {
// 	if len(nums) == 0 || len(nums) == 1 {
// 		return len(nums)
// 	}

// 	// j represents the last non-duplicated character
// 	i, j := 1, 0
// 	for ; i < len(nums); i++ {
// 		if nums[i] != nums[j] {
// 			j++
// 			nums[j] = nums[i]
// 		}
// 	}
// 	j++
// 	fmt.Printf("nums: %+v\n", nums)
// 	return j
// }
