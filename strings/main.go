package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str := os.Args[1:]
	if len(str) == 0 {
		str = []string{"hello"}
	}
	fmt.Println(ReverseText(strings.Join(str, " ")))
}

// ReverseText reverses text within string
func ReverseText(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

// ReverseWords reverses words within a phrase
func ReverseWords(s string, delimiter string) string {
	words := strings.Split(s, delimiter)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, delimiter)
}
