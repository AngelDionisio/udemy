package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMapFunc(t *testing.T) {
	tests := []struct {
		description string
		input       []string
		operation   func(string) string
		expected    []string
	}{
		{
			description: "Single list in array",
			input:       []string{"angel"},
			operation:   UpperCase,
			expected:    []string{"ANGEL"},
		}, {
			description: "Single uppercase difference should return false",
			input:       []string{"Angel", "Luz"},
			operation:   UpperCase,
			expected:    []string{"ANGEL", "LUZ"},
		}, {
			description: "Empty list",
			input:       []string{},
			operation:   UpperCase,
			expected:    nil,
		},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%v, %v", test.description, test.input)
		t.Run(testName, func(t *testing.T) {
			actual := MapFunc(test.input, test.operation)
			if reflect.DeepEqual(actual, test.expected) != true {
				t.Errorf("Expected %v, got %v", test.expected, actual)
			}
		})
	}
}
