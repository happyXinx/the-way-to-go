package main

import "fmt"

func main() {
	printNum(10)
}

func printNum(i int) {
	if i < 1 {
		return
	}
	fmt.Println(i)
	printNum(i - 1)
}
