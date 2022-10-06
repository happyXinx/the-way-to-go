package main

func main() {
	var g int
	go func(i int) {
		s := 0
		print(s)
		for j := 0; j <= i; j++ {
			s += j
		}
		g = s
		print(g)
	}(1000)

	print(g)
}
