package main

import (
	reflect
	testing
)

func TestCleanInput(t *testing.T) {
	case := []struct {
		input	string
		expect	[]string
	}{
		{
			input:		" hello world ",
			expected:	[]string{"hello", "world"}
		},
		{
			input:	"test this string",
			expected:	[]string{"test", "this", "string"}
		},
		{
			input:	"check   this   spaced  out  string  "},
			expected:	[]string{"check", "this", "spaced", "out", "string"}
		}
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		
		if len(actual) != len(c.expected) {
			t.Errorf("Length Error: actual and expected are not equal lengths")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
		}
	}

}
