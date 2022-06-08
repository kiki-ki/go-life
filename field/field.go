package field

import (
	"math/rand"
	"time"
)

const (
	Width  = 100
	Height = 50
)

var (
	Age = 1
)

type FieldType [][]bool

func NewField() FieldType {
	field := make([][]bool, Height)
	for i := range field {
		field[i] = make([]bool, Width)
	}
	return field
}

func NewRandomField() FieldType {
	field := NewField()
	rand.Seed(time.Now().Unix())

	for i := range field {
		for j := range field[i] {
			var cell bool
			if rand.Int()%2 == 1 {
				cell = true
			}
			field[i][j] = cell
		}
	}
	return field
}

func (field FieldType) String() string {
	fStr := ""
	for _, r := range field {
		rStr := ""
		for _, c := range r {
			if c {
				rStr += "*"
			} else {
				rStr += " "
			}
		}
		fStr += rStr + "\n"
	}
	return fStr
}
