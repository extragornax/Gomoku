package gomoku

import "errors"

const (
	boardCellEmpty uint8 = iota
	boardCellOwn         = iota
	boardCellFoe         = iota
)

// Board type represents a board of the gomoku game
type Board struct {
	cells [][]uint8
	size  uint
}

// Init initialises the board
func (brd *Board) init(size uint) error {
	if size > 20 {
		return errors.New(ErrBoardTooLarge)
	} else if size == 0 {
		return errors.New(ErrBoardTooSmall)
	}

	brd.cells = make([][]uint8, size)
	for i := uint(0); i < size; i++ {
		innerLen := uint(i + 1)
		brd.cells[i] = make([]uint8, innerLen)
		for j := uint(0); j < innerLen; j++ {
			brd.cells[i][j] = boardCellEmpty
		}
	}
	brd.size = size
	return nil
}
