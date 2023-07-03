package main

import (
	"fmt"
)

// 发送者
func SendData(ch chan string) {
	fmt.Println("send data")
	ch <- "hello"
	ch <- "hello2"
	ch <- "hello3"
}

// 接收者
func GetData(ch chan string) {
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
//	GetData(ch)
//	go SendData(ch)
//
//	time.Sleep(2e9)
//}
