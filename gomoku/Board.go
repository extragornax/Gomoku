package gomoku

import "errors"

const (
	// BoardCellEmpty is a cell where no one has placed anything
	BoardCellEmpty uint8 = 0
	// BoardCellOwn : your cell
	BoardCellOwn = 1
	// BoardCellFoe : your oponent's cell
	BoardCellFoe = 2
	// BoardCellForbidden : a cell where no one can play
	BoardCellForbidden = 3
)

// Board type represents a board of the gomoku game
type Board struct {
	Cells [][]uint8
	Size  uint
}

// Init initialises the board
func (brd *Board) init(size uint) error {
	if size > 20 {
		return errors.New(ErrBoardTooLarge)
	} else if size == 0 {
		return errors.New(ErrBoardTooSmall)
	}

	brd.Cells = make([][]uint8, size)
	for i := uint(0); i < size; i++ {
		brd.Cells[i] = make([]uint8, size)
		for j := uint(0); j < size; j++ {
			brd.Cells[i][j] = BoardCellEmpty
		}
	}
	brd.Size = size
	return nil
}
