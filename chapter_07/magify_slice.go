package main

import "fmt"

func main() {

	slice := []int{1, 2, 3}
	slice2 := magnifySlice(slice, 3)
	fmt.Println(slice2)
}

func magnifySlice(slice []int, factor int) []int {
	ns := make([]int, len(slice)*factor)
	copy(ns, slice)
	return ns
}
