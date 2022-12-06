package main

import "fmt"

func main() {

	items := make([]map[int]int, 5)

	for i := range items {
		items[i] = make(map[int]int)
		items[i][1] = i
	}
	fmt.Printf("Version A: value of items %v\n", items)

	// Version B: NOT GOOD!
	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		item = make(map[int]int) // item is only a copy of the slice element.
		item[1] = 2              // This 'item' will be lost on the next iteration.
	}
	fmt.Printf("Version B: Value of items: %v\n", items2)
}
