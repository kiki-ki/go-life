package life

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Width  = 200 + 2
	Height = 50 + 2
)

type Field struct {
	Age   int
	field field
}

func NewField() *Field {
	return &Field{
		Age:   1,
		field: newField(),
	}
}

func NewRandomField() *Field {
	f := NewField()
	rand.Seed(time.Now().Unix())

	for y := range f.field {
		for x := range f.field[y] {
			if f.isBufferCell(y, x) {
				continue
			}

			if rand.Intn(2) == 1 {
				f.field[y][x] = true
			}
		}
	}
	return f
}

type field [][]bool

func newField() field {
	f := make([][]bool, Height)
	for y := range f {
		f[y] = make([]bool, Width)
	}
	return f
}

func (f Field) String() string {
	fStr := ""
	for y := range f.field {
		rStr := ""
		for x := range f.field[y] {
			if f.isBufferCell(y, x) {
				continue
			}

			if f.field[y][x] {
				rStr += "○"
			} else {
				rStr += " "
			}
		}
		fStr += rStr + "\n"
	}

	fStr += fmt.Sprintf("\nAge: %d", f.Age)

	return fStr
}

func (f *Field) Next() {
	nextF := newField()

	for y := range f.field {
		for x := range f.field[y] {
			if f.isBufferCell(y, x) {
				continue
			}

			cnt := f.aliveNeighbourhoodCnt(y, x)

			if f.field[y][x] && cnt == 2 || cnt == 3 {
				nextF[y][x] = true
			}
		}
	}

	f.Age++
	f.field = nextF
}

func (f Field) aliveNeighbourhoodCnt(y, x int) int {
	cnt := 0
	for _, c := range f.neighbourhood(y, x) {
		if c {
			cnt++
		}
	}
	return cnt
}

func (f Field) neighbourhood(y, x int) []bool {
	nh := []bool{
		f.field[y-1][x-1], f.field[y-1][x], f.field[y-1][x+1],
		f.field[y][x-1], f.field[y][x+1],
		f.field[y+1][x-1], f.field[y+1][x], f.field[y+1][x+1],
	}
	return nh
}

func (f Field) isBufferCell(y, x int) bool {
	if y == 0 || x == 0 || y == Height-1 || x == Width-1 {
		return true
	}

	return false
}
