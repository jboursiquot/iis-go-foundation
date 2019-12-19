package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var filename string
	if len(os.Args) == 1 {
		filename = "words.txt"
	} else {
		filename = os.Args[1]
	}

	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	counts := wordCount(string(bs))
	fmt.Println(counts)
}

func wordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}
	return m
}
