package main

import "testing"

func TestPalindrome(t *testing.T) {
	tests := []struct {
		description string
		testcase    string
		expected    bool
	}{
		{
			description: "simple string",
			testcase:    "aba",
			expected:    true,
		},
		{
			description: "some special characters",
			testcase:    "A man, a plan, a canal: Panama",
			expected:    true,
		},
		{
			description: "including latin characters",
			testcase:    "ñA man, a plan, a canal: Panamañ",
			expected:    true,
		},
		{
			description: "emojis",
			testcase:    "a😊a",
			expected:    true,
		},
		{
			description: "empty string",
			testcase:    "",
			expected:    true,
		},
		{
			description: "negative case",
			testcase:    "this is a test",
			expected:    false,
		},
	}

	for _, test := range tests {
		actual := IsPalindrome(test.testcase)
		if actual != test.expected {
			t.Errorf("expected  %v to be %v, got: %v",
				test.testcase,
				test.expected,
				actual)
		}
	}
}

func BenchmarkPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}

func BenchmarkPalindromeTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindromeTwo("A man, a plan, a canal: Panama")
	}
}
