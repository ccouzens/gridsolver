package grid

type Square bool

type Clue []int

const (
	White Square = iota
	Black
)

type Grid struct {
	rowClues, columnClues []Clue
	squares               []Square
}
