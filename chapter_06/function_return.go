package main

import "fmt"

func main() {

	f1 := Add2()
	fmt.Println(f1(1))
	f2 := Adder(2)
	fmt.Println(f2(1))
}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
