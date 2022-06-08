package main

import (
	"fmt"
	"go-life/field"
	"time"
)

const (
	span  = 1000
	clear = "\033[2J"
	head  = "\033[1;1H"
)

func main() {
	f := field.NewRandomField()
	fmt.Print(clear)

	for i := 0; i < span; i++ {
		fmt.Print(head)
		fmt.Print(f)
		fmt.Printf("\nAge: %d\n", field.Age)
		f = field.Next(f)
		time.Sleep(100 * time.Millisecond)
	}
}
