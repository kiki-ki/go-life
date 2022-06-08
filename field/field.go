package field

import (
	"math/rand"
	"time"
)

const (
	Width  = 200 + 2
	Height = 60 + 2
)

var (
	Age = 1
)

type FieldType [][]bool

func NewField() FieldType {
	f := make([][]bool, Height)
	for y := range f {
		f[y] = make([]bool, Width)
	}
	return f
}

func NewRandomField() FieldType {
	f := NewField()
	rand.Seed(time.Now().Unix())

	for y := range f {
		for x := range f[y] {
			if isBufferCell(y, x) {
				continue
			}

			var cell bool
			if rand.Int()%2 == 1 {
				cell = true
			}
			f[y][x] = cell
		}
	}
	return f
}

func (f FieldType) String() string {
	fStr := ""
	for y := range f {
		rStr := ""
		for x := range f[y] {
			if isBufferCell(y, x) {
				continue
			}

			if f[y][x] {
				rStr += "○"
			} else {
				rStr += " "
			}
		}
		fStr += rStr + "\n"
	}
	return fStr
}

func isBufferCell(yIdx, xIdx int) bool {
	if yIdx == 0 || xIdx == 0 || yIdx == Height-1 || xIdx == Width-1 {
		return true
	}

	return false
}
