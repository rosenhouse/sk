package sk_test

import (
	"sk"
	"testing"
)

var fullBoardData = []int{
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

func TestValidBoard_Full(t *testing.T) {
	testBoard := sk.BoardFromSlice(fullBoardData)
	if err := testBoard.Validate(); err != nil {
		t.Fatalf("board validation false negative: %s", err)
	}
}

func TestValidBoard_Empty(t *testing.T) {
	testBoard := sk.NewBoard()
	if err := testBoard.Validate(); err != nil {
		t.Fatalf("board validation false negative: %s", err)
	}
}

func TestInvalidBoard_Full_BadRow(t *testing.T) {
	testBoard := sk.BoardFromSlice(fullBoardData)
	testBoard[80] = 9
	err := testBoard.Validate()
	if err == nil {
		t.Fatalf("board validation false positive")
	}
	if err.Error() != "row 9" {
		t.Fatalf("board validation error mismatch: %s", err)
	}
}

func TestInvalidBoard_Simple_BadRow(t *testing.T) {
	testBoard := sk.NewBoard()
	testBoard[0] = 1
	testBoard[8] = 1
	err := testBoard.Validate()
	if err == nil {
		t.Fatalf("board validation false positive")
	}
	if err.Error() != "row 1" {
		t.Fatalf("board validation error mismatch: %s", err)
	}
}

func TestInvalidBoard_Full_BadColumn(t *testing.T) {
	testBoard := sk.BoardFromSlice(fullBoardData)
	testBoard[1] = 9
	testBoard[8] = 0
	err := testBoard.Validate()
	if err == nil {
		t.Fatalf("board validation false positive")
	}
	if err.Error() != "column 2" {
		t.Fatalf("board validation error mismatch: %s", err)
	}
}

func TestInvalidBoard_Simple_BadColumn(t *testing.T) {
	testBoard := sk.NewBoard()
	testBoard[0] = 1
	testBoard[9*8] = 1 // last row, first column
	err := testBoard.Validate()
	if err == nil {
		t.Fatalf("board validation false positive")
	}
	if err.Error() != "column 1" {
		t.Fatalf("board validation error mismatch: %s", err)
	}
}

func TestInvalidBoard_Simple_BadSquare(t *testing.T) {
	testBoard := sk.NewBoard()
	testBoard[0] = 1
	testBoard[10] = 1 // second row, second column
	err := testBoard.Validate()
	if err == nil {
		t.Fatalf("board validation false positive")
	}
	if err.Error() != "square top-left" {
		t.Fatalf("board validation error mismatch: %s", err)
	}

	testBoard = sk.NewBoard()
	testBoard[9*9-1] = 1
	testBoard[8*9-2] = 1 // penultimate row, penultimate column
	err = testBoard.Validate()
	if err == nil {
		t.Fatalf("board validation false positive")
	}
	if err.Error() != "square bottom-right" {
		t.Fatalf("board validation error mismatch: %s", err)
	}
}
