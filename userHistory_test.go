package main

import "testing"

func TestAddIndex(t *testing.T) {
	cases := []struct {
		input    []string
		expected []string
	}{
		{
			input:    []string{"a", "b", "c"},
			expected: []string{"a", "b", "c"},
		},
	}
	for _, c := range cases {
		history := newUserHistory()
		for _, v := range c.input {
			history.add(v)
		}
		for i := range c.expected {
			if history.history[i] != c.expected[i] {
				t.Fail()
				return
			}
		}
	}
}
func TestIncrementIndex(t *testing.T) {
	cases := []string{"a", "b", "c"}
	expectedDec := []string{"c", "b", "a", "", ""}
	expectedIn := []string{"b", "c", "", ""}
	history := newUserHistory()
	for _, v := range cases {
		history.add(v)
	}
	for _, v := range expectedDec {
		element, _ := history.decrementIndex()
		if v != element {
			t.Errorf("decrement Index - actual: %v != expected: %v", element, v)
		}
	}
	for _, v := range expectedIn {
		element, _ := history.incrementIndex()
		if v != element {
			t.Errorf("increment Index - actual: %v != expected: %v", element, v)
		}
	}

}
