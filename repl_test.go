package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   Hello   WORLD   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "123 ==== !ko (LoL)",
			expected: []string{"123", "====", "!ko", "(lol)"},
		},
		{
			input:    "",
			expected: []string{"", ""},
		},
		{
			input:    "%%% \n",
			expected: []string{"%%%", ""},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf(
				`test failed; actual output does not match the length of expected:
			Actual: %v
			Expected: %v`,
				actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("test failed; word('%v') does not match expected('%v')\n", word, expectedWord)
			}
		}
	}
}
