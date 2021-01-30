package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/angeldionisio/udemy/golang/random/concurrency/13_xkcd_exercise/xkcd"
)

func main() {
	comicsToFetch := generateSliceOfRandomNumbers(10, 2380)

	comics := xkcd.GetComicsAsync(comicsToFetch)
	log.Println(comics)
}

func generateSliceOfRandomNumbers(size, maxNum int) []int {
	rand.Seed(time.Now().UnixNano())
	list := make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = rand.Intn(maxNum)
	}
	return list
}
