package main

func sum(x, y int, c chan int) {
	c <- x + y
}

//
//func main() {
//	c := make(chan int)
//	go sum(1, 2, c)
//	print(<-c)
//}
