package main

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name                string
	nonExportedAgeField string
}

func main() {
	t := template.New("1")
	t, _ = t.Parse("hello {{.nonExportedAgeField}}")
	p := Person{
		Name:                "Mary",
		nonExportedAgeField: "21",
	}
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}
