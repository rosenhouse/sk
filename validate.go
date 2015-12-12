package sk

import (
	"bytes"
	"fmt"
)

type ValidationError struct {
	Errors []string
}

func (e *ValidationError) addRowError(r int) {
	e.Errors = append(e.Errors, fmt.Sprintf("row %d", r+1))
}

func (e *ValidationError) addColumnError(c int) {
	e.Errors = append(e.Errors, fmt.Sprintf("column %d", c+1))
}

func (e *ValidationError) addSquareError(r, c int) {
	loc := ""
	e.Errors = append(e.Errors, fmt.Sprintf("square %s", loc))
}

func newValidationError(s ...string) *ValidationError {
	return &ValidationError{Errors: s}
}

func (e *ValidationError) Error() string {
	if e == nil {
		return ""
	}
	all := e.Errors
	b := &bytes.Buffer{}
	for i := 0; i < len(all); i++ {
		b.WriteString(all[i])
		if i+1 < len(all) {
			b.WriteString("; ")
		}
	}
	return b.String()
}

func (b Board) Validate() error {
	err := &ValidationError{}
	for i := 0; i < 9; i++ {
		if !b.isValidRow(i) {
			err.addRowError(i)
			return err
		}
		if !b.isValidColumn(i) {
			err.addColumnError(i)
			return err
		}
	}
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if !b.isValidSquare(r, c) {
				err.addSquareError(r, c)
				return err
			}
		}
	}

	return nil
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
