package main

import "fmt"

func main() {
	sl := []byte{'a', 'u', 'z', 's'}
	data := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	sliceVal := Append(sl, data)
	fmt.Printf("slice = %c \n", sliceVal)
}

func Append(slice, data []byte) []byte {

	// append的实现原理
	lenNewSlice := len(slice) + len(data)
	capNewSlice := cap(slice)
	if cap(slice) < lenNewSlice {
		capNewSlice = lenNewSlice
	}

	newSlice := make([]byte, lenNewSlice, capNewSlice)
	for key, item := range slice {
		newSlice[key] = item
	}
	for key, item := range data {
		newSlice[key+len(slice)] = item
	}
	return newSlice
}
