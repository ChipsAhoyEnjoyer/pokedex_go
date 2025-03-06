package main

import "strings"

func cleanInput(text string) []string {
	t := strings.ToLower(text)
	s := strings.Fields(t)
	return s
}
