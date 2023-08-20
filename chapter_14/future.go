package main

type Matrix struct {
}

func InverseProduct(a Matrix, b Matrix) Matrix {
	a_inv_future := InverseFuture(a) // start as a goroutine
	b_inv_future := InverseFuture(b) // start as a goroutine
	a_inv := <-a_inv_future
	b_inv := <-b_inv_future
	return Product(a_inv, b_inv)
}

func InverseFuture(a Matrix) chan Matrix {
	future := make(chan Matrix)
	go func() {
		future <- Inverse(a)
	}()
	return future
}

func InverseFuture(a Matrix) chan Matrix {
	future := make(chan Matrix)
	go func() {
		future <- Inverse(a)
	}()
	return future
}
