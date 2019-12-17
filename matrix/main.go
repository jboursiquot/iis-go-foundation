package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {

	add := func(x, y int) int {
		return x + y
	}

	div := func(x, y int) int {
		return int(x / y)
	}

	type op func(int, int) int

	ops := make(map[string]op, 2)
	ops["add"] = add
	ops["div"] = div
	ops["sub"] = sub

	// fmt.Println(ops["add"](1, 2))
	// fmt.Println(ops["div"](3, 2))
	// fmt.Println(ops["sub"](4, 3))

	nums := [][]int{
		[]int{1, 2},
		[]int{3, 2},
		[]int{4, 3},
	}
	sOps := []op{add, div, sub}
	for _, pair := range nums {
		for _, op := range sOps {
			r := op(pair[0], pair[1])
			// fmt.Printf("Pair: %v, Operation: %v, Result: %v\n", pair, op, r)
			// fmt.Printf("Pair: %v, Operation: %T, Result: %v\n", pair, op, r)

			fn := runtime.FuncForPC(reflect.ValueOf(op).Pointer()).Name()
			fmt.Printf("Pair: %v, Operation: %v, Result: %v\n", pair, fn, r)
		}
	}
}

func sub(x, y int) int {
	return x - y
}
