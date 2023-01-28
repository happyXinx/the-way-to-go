package main

//
//func main() {
//	ch1 := make(chan int)
//	go pump(ch1)
//	println(<-ch1)
//	println(<-ch1)
//}

func pump(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}
