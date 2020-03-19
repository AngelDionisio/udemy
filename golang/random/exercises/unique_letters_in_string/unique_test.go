package main

import (
	"testing"
)

func TestAbs(t *testing.T) {
	var tests = []struct {
		testCase   string
		testString string
		expected   bool
	}{
		{
			testCase:   "no repeated characters",
			testString: "abcd",
			expected:   true,
		},
		{
			testCase:   "repated characters",
			testString: "aofo",
			expected:   false,
		},
		{
			testCase:   "empty string",
			testString: "",
			expected:   true,
		},
		{
			testCase:   "single character",
			testString: "a",
			expected:   true,
		},
	}

	for _, test := range tests {
		result := Unique(test.testString)
		if result != test.expected {
			t.Errorf("expected %v to yield %v, got: %v", test.testString, test.expected, result)
		}
	}

}

func BenchmarkUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Unique("angel")
	}
}
