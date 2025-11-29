package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input	string
		expected	[]string
	}{
		{
			input:		" hello world ",
			expected:	[]string{"hello", "world"},
		},
		{
			input:	"test this string",
			expected:	[]string{"test", "this", "string"},
		},
		{
			input:	"check   this   spaced  out  string  ",
			expected:	[]string{"check", "this", "spaced", "out", "string"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		
		actualLen := len(actual)
		expectedLen := len(c.expected)

		if actualLen != expectedLen {
			t.Errorf("Length Error: got %v, want %v", actualLen, expectedLen)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			
			if word != expectedWord {
				t.Errorf("Match Error: got %v, want %v", word, expectedWord)
				t.Fatal()
			}
		}
	}

	if t.Failed() {
		t.Fatal("Test Failed! See errors above")
	}

}
