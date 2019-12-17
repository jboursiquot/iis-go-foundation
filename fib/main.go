package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	start, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	end, _ := strconv.Atoi(os.Args[2])
	seq := genFibSequence(start, end)
	fmt.Println(seq)
}

func genFibSequence(start, end int) []int {
	result := []int{}
	for n := start; n <= end; n++ {
		result = append(result, fib(n))
	}
	return result
}

// recursive solutions - there are others
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
