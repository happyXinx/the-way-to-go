package main

func main() {

	var ar = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := ar[0:0]
	s2 := ar[0:1]
	println(len(s1))
	println(len(s2))

}
