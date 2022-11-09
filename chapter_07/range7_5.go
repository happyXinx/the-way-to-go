package main

func main() {

	items := [...]int{10, 20, 30, 40, 50}

	for _, item := range items {
		item *= 2
	}

	for i := 0; i < len(items); i++ {
		items[i] *= 2
	}

	for _, item := range items {
		println(item)
	}
}
