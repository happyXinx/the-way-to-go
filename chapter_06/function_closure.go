package main

import "fmt"

func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}

func main() {
	f := Adder()
	fmt.Print(f(1), "--")
	fmt.Print(f(10), "--")
	fmt.Println(f(100), "--")
}
