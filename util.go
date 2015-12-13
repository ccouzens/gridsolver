package grid

func (g *Grid) SquareAt(x, y int) Square {
	width := len(g.rowClues)
	return g.squares[x+y*width]
}
