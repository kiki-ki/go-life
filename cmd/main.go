package main

import (
	"fmt"
	"go-life/life"
	"time"
)

const (
	span  = 1000
	clear = "\033[2J"
	head  = "\033[1;1H"
)

func main() {
	field := life.NewRandomField()
	fmt.Print(clear)

	for i := 0; i < span; i++ {
		fmt.Print(head)
		fmt.Print(field)
		field.Next()
		time.Sleep(100 * time.Millisecond)
	}
}
