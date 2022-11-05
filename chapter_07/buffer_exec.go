package main

import (
	"bytes"
	"fmt"
)

func main() {

	var buffer bytes.Buffer
	buffer.WriteString("1")
	buffer.WriteString("2")
	fmt.Println(buffer.String())

}

