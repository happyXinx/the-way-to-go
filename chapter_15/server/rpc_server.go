package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"the-way-to-go/chapter_15/rpc_objects"
	"time"
)

func main() {
	calc := new(rpc_objects.Args)
	rpc.Register(calc)
	rpc.HandleHTTP()
	listener, e := net.Listen("tcp", "localhost:1234")
	if e != nil {
		log.Fatal("Starting RPC-server -listen error:", e)
	}

	// 对于每一个进入lisntener的请求，都是由协程去启用一个http.Serve(listener, nil)
	// , 为每一个传入HTtP连接创建一个新的服务线程。
	go http.Serve(listener, nil)
	time.Sleep(1000e9)
}

/* Output:
Starting Process E:/Go/GoBoek/code_examples/chapter_14/rpc_server.exe ...

** after 5 s: **
End Process exit status 0
*/
