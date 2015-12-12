package sk

func (b Board) IsValid() bool {
	ok := true
	for i := 0; i < 9; i++ {
		ok = ok && b.isValidRow(i)
		ok = ok && b.isValidColumn(i)
	}
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			ok = ok && b.isValidSquare(r, c)
		}
	}

	return ok
}

func isValidSet(toValidate []int) bool {
	if len(toValidate) != 9 {
		panic("unexpected length")
	}
	histogram := [9]int{}
	for i := 0; i < 9; i++ {
		if toValidate[i] == 0 {
			continue
		}
		hi := toValidate[i] - 1
		histogram[hi]++
		if histogram[hi] > 1 {
			return false
		}
	}
	return true
}

func (b Board) isValidRow(r int) bool {
	rowData := b[r*9 : (r+1)*9]
	return isValidSet(rowData)
}

func (b Board) isValidColumn(c int) bool {
	columnData := make([]int, 9)
	for i := 0; i < 9; i++ {
		columnData[i] = b[i*9+c]
	}
	return isValidSet(columnData)
}

func (b Board) isValidSquare(r, c int) bool {
	squareData := make([]int, 9)
	for rr := 0; rr < 3; rr++ {
		for cc := 0; cc < 3; cc++ {
			i := 3*rr + cc
			fullRow := 3*r + rr
			fullColumn := 3*c + cc
			squareData[i] = b[9*fullRow+fullColumn]
		}
	}
	return isValidSet(squareData)
}
