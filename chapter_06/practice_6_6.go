package main

func main() {
	print(jieCheng(4))
}

func jieCheng(n int) int {
	if n <= 1 {
		return 1
	}
	return n * jieCheng(n-1)
}
