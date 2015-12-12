package sk_test

import (
	"sk"
	"testing"
)

var displayTestData = []int{
	1, 2, 3, 4, 5, 6, 7, 8, 9,
	4, 5, 6, 7, 8, 9, 1, 2, 3,
	7, 8, 9, 1, 2, 3, 4, 5, 6,
	2, 3, 4, 5, 6, 7, 8, 9, 1,
	5, 6, 7, 8, 9, 1, 2, 3, 4,
	8, 9, 1, 2, 3, 4, 5, 6, 7,
	3, 4, 5, 6, 7, 8, 9, 1, 2,
	6, 7, 8, 9, 1, 2, 3, 4, 5,
	9, 1, 2, 3, 4, 5, 6, 7, 8,
}

func TestDisplayPrintBoard(t *testing.T) {
	testBoard := sk.BoardFromSlice(displayTestData)
	actual, err := sk.Display(testBoard)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	expected := `
1 2 3 | 4 5 6 | 7 8 9
4 5 6 | 7 8 9 | 1 2 3
7 8 9 | 1 2 3 | 4 5 6
- - - + - - - + - - -
2 3 4 | 5 6 7 | 8 9 1
5 6 7 | 8 9 1 | 2 3 4
8 9 1 | 2 3 4 | 5 6 7
- - - + - - - + - - -
3 4 5 | 6 7 8 | 9 1 2
6 7 8 | 9 1 2 | 3 4 5
9 1 2 | 3 4 5 | 6 7 8
`
	if actual != expected {
		t.Fatalf(
			"wrong board: \nexpected:\n%s\nactual:\n%s",
			expected, actual,
		)
	}
}

func TestDisplayEmptyBoard(t *testing.T) {
	testBoard := sk.BoardFromSlice(make([]int, 9*9))
	actual, err := sk.Display(testBoard)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	expected := `
      |       |      
      |       |      
      |       |      
- - - + - - - + - - -
      |       |      
      |       |      
      |       |      
- - - + - - - + - - -
      |       |      
      |       |      
      |       |      
`
	if actual != expected {
		t.Fatalf(
			"wrong board: \nexpected:\n%s\nactual:\n%s",
			expected, actual,
		)
	}
}
