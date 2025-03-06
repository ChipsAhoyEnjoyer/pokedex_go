package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	s := strings.Fields(text)
	for i := range s {
		s[i] = strings.ToLower(s[i])
	}
	return s
}
