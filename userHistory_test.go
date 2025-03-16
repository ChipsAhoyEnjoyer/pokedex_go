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
	cases := []struct {
		input    []string
		expected []string
	}{
		{
			input:    []string{"a", "b", "c"},
			expected: []string{"", "", ""},
		},
		{
			input:    []string{},
			expected: []string{""},
		},
	}
	for _, c := range cases {
		history := newUserHistory()
		for _, v := range c.input {
			history.add(v)
		}
		for i := range c.expected {
			actual, _ := history.incrementIndex()
			if actual != c.expected[i] {
				t.Fail()
				return
			}
		}
		if history.historyLen != len(c.input) || history.currentIndex > len(c.input)-1 {
			t.Fail()
			return
		}
	}
}
func TestDecrementIndex(t *testing.T) {
	cases := []struct {
		input    []string
		expected []string
	}{
		{
			input:    []string{"a", "b", "c"},
			expected: []string{"c", "b", "a"},
		},
		{
			input:    []string{},
			expected: []string{""},
		},
	}
	for _, c := range cases {
		history := newUserHistory()
		for _, v := range c.input {
			history.add(v)
		}
		for i := range c.expected {
			actual, _ := history.decrementIndex()
			if actual != c.expected[i] {
				t.Fail()
				return
			}
		}
		if history.historyLen != len(c.input) || history.currentIndex > 0 {
			t.Fail()
			return
		}
	}
}
