package main

import (
	"fmt"
	"time"
)

func main() {

	stamp := int64(1584109800000)
	x := isInvalidStartTime(stamp)

	fmt.Println(x)
}

func isInvalidStartTime(s int64) bool {
	requestedDate := time.Unix(0, s*int64(time.Millisecond)).UTC()

	// remove timezone from time, this allows parsing the string
	// in UTC time, this allows use of time.Now() regardless of
	// where the machine's timezone is
	timeString := time.Now().Format(time.RFC3339)[:19]

	parsed, err := time.Parse("2006-01-02T15:04:05", timeString)
	if err != nil {
		fmt.Println("ERRRRROOOORRRRRRR", err)
	}

	// add timezone info to allow comparison
	today := parsed.UTC()
	fmt.Println("requested:", requestedDate)
	fmt.Println("now      :", today)

	return requestedDate.Before(today)
}

func isInvalidStartTimeOLD(s int64) bool {
	startTime := s / 1000
	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc).Unix()

	today := time.Unix(now, 0).In(loc)
	requestedDate := time.Unix(startTime, 0).In(loc)

	fmt.Println("now:      ", now)
	fmt.Println("today:    ", today)
	fmt.Println("requested:", requestedDate)

	return requestedDate.Before(today)
}
