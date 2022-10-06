package main

import "time"

func main() {
	start := time.Now()
	f := fibonacciClosure()
	for i := 0; i <= 40; i++ {
		println(f())
	}
	end := time.Now()
	println(end.Sub(start))
}

func fibonacciClosure() func() int {
	f1, f2 := 1, 1
	return func() int {
		f1, f2 = f2, f1+f2
		return f1 + f2
	}
}
