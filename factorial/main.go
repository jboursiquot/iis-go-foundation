package main

import (
	"fmt"
	"os"
	"strconv"
)

func factorial(n int) (result int) { // result var predefined
	if n < 0 || n > 100 {
		panic("MISBEHAVING!")
	}
	if n > 0 {
		result = n * factorial(n-1)
		return
	}
	return 1
}

func main() {
	if n, err := strconv.Atoi(os.Args[1]); err == nil {
		fmt.Println(factorial(n))
		os.Exit(0)
	}
	panic("BOOM!")
}
