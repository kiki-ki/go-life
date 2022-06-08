package field

func Next(f FieldType) FieldType {
	nextF := NewField()

	for y := range f {
		for x := range f[y] {
			if isBufferCell(y, x) {
				continue
			}

			nh := []bool{
				f[y-1][x-1], f[y-1][x], f[y-1][x+1],
				f[y][x-1], f[y][x+1],
				f[y+1][x-1], f[y+1][x], f[y+1][x+1],
			}
			aliveNhCnt := aliveNeighbourhoodCnt(nh)

			if f[y][x] {
				if aliveNhCnt <= 1 || aliveNhCnt >= 4 {
					nextF[y][x] = false
				}
			} else {
				if aliveNhCnt == 3 {
					nextF[y][x] = true
				}
			}
		}
	}

	Age++

	return nextF
}

func aliveNeighbourhoodCnt(nh []bool) int {
	cnt := 0
	for _, c := range nh {
		if c {
			cnt++
		}
	}
	return cnt
}
