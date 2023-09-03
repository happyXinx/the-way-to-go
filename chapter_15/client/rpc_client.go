package main

import (
	"fmt"
	"log"
	"net/rpc"
	"the-way-to-go/chapter_15/rpc_objects"
)

const serverAddress = "localhost"

func main() {
	clent, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("Error dialing:", err)
	}
	args := &rpc_objects.Args{7, 8}
	var reply int
	// 客户端必须知道服务器端定义的类型和它的方法
	err = clent.Call("Args.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Args error:", err)
	}
	fmt.Printf("Args: %d * %d = %d", args.N, args.M, reply)
}
