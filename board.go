package sk

type Board []int

func BoardFromSlice(b []int) Board {
	return Board(b)
}

func (b Board) IsValid() bool {
	return true
}
