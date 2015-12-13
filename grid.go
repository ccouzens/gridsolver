package grid

type Square bool

type Clue []int

const (
	White Square = false
	Black Square = true
)

type Grid struct {
	rowClues, columnClues []Clue
	squares               []Square
}
