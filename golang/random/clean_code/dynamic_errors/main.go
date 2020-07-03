package main

import (
	"errors"
	"fmt"
)

// Returning Dynamic Errors

// ErrItemNotFound error for items not found in a map
var (
	ErrItemNotFound = errors.New("item could not be found in the store")
	NullInt         = -1
	commits         = map[string]int{
		"codeLord": 7894,
		"fooBar":   74515,
		"otacon":   456785,
	}
)

// ErrorDetails provides a way to express errors with custom details
type ErrorDetails interface {
	Error() string
	Type() error
}

type errDetails struct {
	errType error
	details interface{}
}

// NewErrorDetails constructor
func NewErrorDetails(err error, details ...interface{}) ErrorDetails {
	return &errDetails{
		errType: err,
		details: details,
	}
}

func (err *errDetails) Error() string {
	return fmt.Sprintf("%v, %v", err.errType, err.details)
}

func (err *errDetails) Type() error {
	return err.errType
}

func main() {
	handle := "blahbleh"
	c, err := GetCommitsByHandle(handle)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("# of commits for %v: %d", handle, c)
}

// GetCommitsByHandle finds items in commit map
func GetCommitsByHandle(handle string) (int, error) {
	c, ok := commits[handle]
	if !ok {
		return NullInt, NewErrorDetails(
			ErrItemNotFound,
			fmt.Sprintf("could not find item with handle: %v", handle),
		)
	}

	return c, nil
}
