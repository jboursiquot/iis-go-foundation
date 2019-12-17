package main

import "fmt"

func main() {
	funcA()
	fmt.Println("finishing...")
}
func funcC() {
	fmt.Println("funcC")
	panic("Error!")
}
func funcB() {
	defer defer1()
	funcC()
}

func funcA() {
	defer defer2()
	funcB()
}
func defer1() {
	fmt.Println("defer1")
}
func defer2() {
	s := recover()
	fmt.Println("recovered from", s)
}
