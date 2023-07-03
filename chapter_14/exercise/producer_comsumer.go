package main

func producer(c chan int) {
	for i := 0; i < 100; i += 10 {
		c <- i
	}
	close(c)
}

func consumer(c chan int, c1 chan bool) {
	for num := range c {
		println(num)
	}
	c1 <- true
}

func main() {
	c := make(chan int)
	c1 := make(chan bool)
	go producer(c)
	go consumer(c, c1)
	<-c1
}
