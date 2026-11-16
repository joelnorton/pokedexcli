package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " Here we go again",
			expected: []string{"here", "we", "go", "again"},
		},
		{
			input:    "It's the end of the world as we know it",
			expected: []string{"it's", "the", "end", "of", "the", "world", "as", "we", "know", "it"},
		},
		{
			input:    "I feel fine",
			expected: []string{"i", "feel", "fine"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice
		// if they don't match, use t.Errorf and continue to the next case
		if len(actual) != len(c.expected) {
			t.Errorf("Lengths do not match\nExpected: %d\nReturned:%d\n", len(c.expected), len(actual))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Words do not match\nExpected: %s\nReturned:%s\n", expectedWord, word)
				return
			}
		}
	}
}
