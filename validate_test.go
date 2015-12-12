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

func TestValidBoard(t *testing.T) {
	testBoard := sk.BoardFromSlice(fullBoardData)
	if !testBoard.IsValid() {
		t.Fatalf("board validation false negative")
	}
}

func TestValidBoard_Empty(t *testing.T) {
	testBoard := sk.NewBoard()
	if !testBoard.IsValid() {
		t.Fatalf("board validation false negative")
	}
}

func TestInvalidBoard_Full_BadRow(t *testing.T) {
	testBoard := sk.BoardFromSlice(fullBoardData)
	testBoard[80] = 9
	if testBoard.IsValid() {
		t.Fatalf("board validation false positive")
	}
}

func TestInvalidBoard_Simple_BadRow(t *testing.T) {
	testBoard := sk.NewBoard()
	testBoard[0] = 1
	testBoard[8] = 1
	if testBoard.IsValid() {
		t.Fatalf("board validation false positive")
	}
}

func TestInvalidBoard_Full_BadColumn(t *testing.T) {
	testBoard := sk.BoardFromSlice(fullBoardData)
	testBoard[1] = 9
	testBoard[8] = 0
	if testBoard.IsValid() {
		t.Fatalf("board validation false positive")
	}
}

func TestInvalidBoard_Simple_BadColumn(t *testing.T) {
	testBoard := sk.NewBoard()
	testBoard[0] = 1
	testBoard[9*8] = 1 // last row, first column
	if testBoard.IsValid() {
		t.Fatalf("board validation false positive")
	}
}

func TestInvalidBoard_Simple_BadSquare(t *testing.T) {
	testBoard := sk.NewBoard()
	testBoard[0] = 1
	testBoard[10] = 1 // second row, second column
	if testBoard.IsValid() {
		t.Fatalf("board validation false positive")
	}

	testBoard = sk.NewBoard()
	testBoard[9*9-1] = 1
	testBoard[8*9-2] = 1 // penultimate row, penultimate column
	if testBoard.IsValid() {
		t.Fatalf("board validation false positive")
	}
}
