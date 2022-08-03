package main

import (
	"fmt"
	"testing"
)

var table = []struct {
	numList []int
}{
	{numList: []int{1, 2, 3, 4, 5}},
	{numList: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{numList: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}},
}

func BenchmarkBinarySearch(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input size: [%v]", len(v.numList)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				BinarySearch(v.numList, 5)
			}
		})
	}
}
