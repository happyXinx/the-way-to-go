package main

import (
	"fmt"
)

func sendData(ch chan string) {
	fmt.Println("send data")
	ch <- "hello"
	ch <- "hello2"
	ch <- "hello3"
}

func getData(ch chan string) {
	var input string
	fmt.Println("get data")
	for {
		input = <-ch
		fmt.Println(input)
	}
}

//
//func main() {
//	ch := make(chan string)
//	go getData(ch)
//	go sendData(ch)
//
//	time.Sleep(2e9)
//}
