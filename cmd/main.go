package main

import (
	"fmt"
	"go-life/field"
)

func main() {
	f := field.NewRandomField()

	for i := 0; i < 100; i++ {
		fmt.Println(f)
		f = field.Next(f)
	}
}
