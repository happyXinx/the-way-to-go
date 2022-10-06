package main

import "strings"

func MakeAddSuffix(suffix string) func(name string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	addBmp := MakeAddSuffix(".bmp")
	println(addBmp("aaa"))
}
