package main

import "fmt"

func f1(in chan int) {
	fmt.Println("a")
	fmt.Println(<-in)
}

func main() {
	var ch chan int
	//go f1(ch)
	ch <- 2
	println(<-ch)
}
