package main

type Request struct {
	a, b   int
	replyc chan int
}

type binOp func(a, b int) int

func run(op binOp, req *Request) {
	req.replyc <- op(req.a, req.b)
}

func server(op binOp, service chan *Request, quit chan bool) {
	for {
		select {
		case req := <-service:
			go run(op, req)
		case <-quit:
			println("quit")
			return
		}
	}
}

func startServer(op binOp) (service chan *Request, quit chan bool) {
	service = make(chan *Request)
	quit = make(chan bool)
	go server(op, service, quit)
	return service, quit
}

//
//func main() {
//
//	adder, quit := startServer(func(a, b int) int {
//		return a + b
//	})
//	const N = 100
//
//	var reqs [N]Request
//
//	for i := 0; i < N; i++ {
//		req := &reqs[i]
//		req.a = i
//		req.b = i + N
//		req.replyc = make(chan int)
//		adder <- req
//	}
//	quit <- true
//	for i := N - 1; i >= 0; i-- { // 顺序无所谓
//		if <-reqs[i].replyc != N+2*i {
//			fmt.Println("fail at", i)
//		} else {
//			fmt.Println("Request ", i, "is ok!")
//		}
//	}
//
//	fmt.Println("done")
//}
