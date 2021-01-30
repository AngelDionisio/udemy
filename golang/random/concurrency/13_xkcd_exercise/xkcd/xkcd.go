package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const baseXkcdURL = "https://xkcd.com/%d/info.0.json"

// Comic represents XKCD comic JSON data
type Comic struct {
	Num   int    `json:"num"`
	Link  string `json:"link"`
	Img   string `json:"img"`
	Title string `json:"title"`
}

// generateURL using base XKCD's base URL
func generateURL(comicID int) string {
	return fmt.Sprintf(baseXkcdURL, comicID)
}

// GetComic fetches an XKCD comic, communicates result via channel.
func GetComic(comicID int, c chan Comic, wg *sync.WaitGroup) error {
	defer wg.Done()

	url := generateURL(comicID)
	resp, err := http.Get(url)
	if err != nil {
		c <- Comic{Num: comicID}
	}
	defer resp.Body.Close()

	var comic Comic
	err = json.NewDecoder(resp.Body).Decode(&comic)
	if err != nil {
		return fmt.Errorf("Error decoding comicID: %v", comicID)
	}

	c <- comic

	return nil
}

// GetComicsAsync gets a list of XKCD comics async.
func GetComicsAsync(comicIDs []int) map[int]Comic {
	start := time.Now()
	defer func() {
		fmt.Printf("Channels fetches took: %v\n", time.Since(start))
	}()

	var wg sync.WaitGroup
	c := make(chan Comic)

	// wait for all jobs in WaitGroup to complete, then close channels
	go func() {
		wg.Wait()
		close(c)
	}()

	for _, comidID := range comicIDs {
		wg.Add(1)
		go GetComic(comidID, c, &wg)
	}

	m := make(map[int]Comic)
	for comic := range c {
		m[comic.Num] = comic
	}

	return m

}

// Worker worker pool for comics
func Worker(id int, wg *sync.WaitGroup, jobs <-chan int, results chan<- Comic, errors chan<- error) {
	defer wg.Done()

	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		c, err := getComicPlain(j)
		if err != nil {
			errors <- fmt.Errorf("Failed getting comicID: %v due to %v", j, err)
		}

		results <- *c
	}
}

// getComic gets a comic given an ID
func getComicPlain(comicID int) (comic *Comic, err error) {
	url := generateURL(comicID)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(response.Body).Decode(&comic)
	if err != nil {
		return nil, err
	}

	return comic, nil
}
