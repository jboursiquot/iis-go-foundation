package main

import (
	"fmt"
	"os"
	"strconv"
)

func factorial(n int64) (result int64) { // result var predefined
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
	defer func() {
		err := recover()
		fmt.Printf("recover from %s", err)
	}()

	if n, err := strconv.Atoi(os.Args[1]); err == nil {
		m := int64(n)
		fmt.Println(factorial(m))
		os.Exit(0)
	}
	panic("BOOM!")
}
