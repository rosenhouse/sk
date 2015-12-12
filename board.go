package sk

type Board []int

func NewBoard() Board {
	return Board(make([]int, 9*9))
}

func BoardFromSlice(b []int) Board {
	myCopy := make([]int, 9*9)
	copy(myCopy, b)
	return Board(myCopy)
}
