package main

import (
	"fmt"
)

func f1(in chan int) {
	fmt.Println("a")
	fmt.Println(<-in)
}

//
//func main() {
//	ch := make(chan int)
//	ch <- 2 // 这里一直阻塞，运行不到下面
//	go f1(ch)
//
//}
