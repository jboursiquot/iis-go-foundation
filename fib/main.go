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
	seq := fibBetween(start, end)
	fmt.Println(seq)
}

func fibBetween(min, max int) []int {
	result := []int{}
	n1, n2, fib := 1, 1, 0

	for {
		if n1+n2 < min {
			fib = n1 + n2
			n1 = n2
			n2 = fib
			continue
		}
		break
	}

	for {
		if n1+n2 <= max {
			fib = n1 + n2
			result = append(result, fib)
			n1 = n2
			n2 = fib
			continue
		}
		break
	}

	return result
}
