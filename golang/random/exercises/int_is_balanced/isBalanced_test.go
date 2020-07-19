package main

import (
	"fmt"
	"testing"
)

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expected    bool
	}{
		{
			description: "open parens",
			input:       "{}",
			expected:    true,
		}, {
			description: "empty string",
			input:       "",
			expected:    false,
		}, {
			description: "multiple nested parens",
			input:       "[({abc})]",
			expected:    true,
		}, {
			description: "not balanced",
			input:       "[({abc}}})]",
			expected:    false,
		},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%v, %v", test.description, test.input)
		t.Run(testName, func(t *testing.T) {
			actual := IsBalanced(test.input)
			if actual != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, actual)
			}
		})
	}
}
