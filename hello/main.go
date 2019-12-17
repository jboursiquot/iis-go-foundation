package main

import (
	"fmt"
)

func main() {
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}
	if array1 == array2 {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
	// slice1 := []int{1, 2, 3}
	// slice2 := []int{1, 2, 3}
	// if slice1 == slice2 {
	// 	fmt.Println("equal")
	// } else {
	// 	fmt.Println("not equal")
	// }
}
