package main

import (
	"os"
	"strings"
	"testing"
)

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
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "   ",
			expected: []string{},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Lengths do not match\nExpected: %d\nReturned:%d\n", len(c.expected), len(actual))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Words do not match\nExpected: %s\nReturned:%s\n", expectedWord, word)
				continue
			}
		}
	}
}

func TestStartRepl_EmptyInputNoPanic(t *testing.T) {
	// Feed a blank line followed by EOF via os.Stdin
	f, err := os.CreateTemp("", "stdin")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if _, err := f.WriteString("\n"); err != nil {
		t.Fatal(err)
	}
	if _, err := f.Seek(0, 0); err != nil {
		t.Fatal(err)
	}

	old := os.Stdin
	os.Stdin = f
	defer func() { os.Stdin = old }()

	out := captureOutput(func() {
		startRepl(&config{
			cliCommands: map[string]cliCommand{
				"exit": {"exit", "Exit", commandExit},
			},
		})
	})

	if !strings.Contains(out, "Pokedex >") {
		t.Errorf("expected prompt in output, got %q", out)
	}
}
