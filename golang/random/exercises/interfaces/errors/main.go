package main

import (
	"fmt"
	"time"
)

// MyError captures error reason and time
type MyError struct {
	When   time.Time
	Reason string
}

// Error, now MyError implements the error interface
func (e *MyError) Error() string {
	return fmt.Sprintf("Error: at %v, due to %s:", e.When, e.Reason)
}

func run() error {
	return &MyError{
		When:   time.Now(),
		Reason: "user generated error",
	}
}

// ErrNegativeSqrt to allow custom error for square root operations
// because it implements the Error interface, this type can also be used as
// type error
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g.", float64(e))
}

// Sqrt returns square root of postitive numbers
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0
	for i := 0; i < 10; i++ {
		z = z - ((z*z)-x)/(2*z)
	}
	return z, nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
