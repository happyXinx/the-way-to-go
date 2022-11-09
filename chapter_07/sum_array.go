package main

import "fmt"

func main() {
	var a_array = [4]float32{1.0, 2.0, 3.0, 4.0}
	var a_slice = []float32{1.0, 2.0, 3.0, 4.0}
	fmt.Printf("The sum of the array is: %f\n", Sum(a_array[:]))
	fmt.Printf("The sum of the array is: %f\n", Sum(a_slice))
}

func SumArray(a [4]float32) (sum float32) {
	for _, item := range a {
		sum += item
	}
	return
}

func Sum(a []float32) (sum float32) {
	for _, item := range a {
		sum += item
	}
	return
}

func SumAndAverage(a []int) (int, float32) {
	sum := 0
	for _, item := range a {
		sum += item
	}
	return sum, float32(sum / len(a))
}
