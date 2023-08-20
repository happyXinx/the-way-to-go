package main

const (
	AvailableMemory = 10 << 20

	// 10 MB, 示例

	AverageMemoryPerRequest = 10 << 10

	// 10 KB

	MAXREQS = AvailableMemory / AverageMemoryPerRequest
	// 原文中说 MAXREQS 是 1000，实际计算是 1024 ，后面按照原文的 1000 来描述

)

func handler(r *Request) {
	
}
