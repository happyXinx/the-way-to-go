package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	result := 0
	for i := 0; i <= 40; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	println(end.Sub(start))
}

func fibonacci(i int) int {
	if i <= 1 {
		return 1
	}
	return fibonacci(i-1) + fibonacci(i-2)
}
