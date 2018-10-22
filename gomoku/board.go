package gomoku

import "errors"

const (
	boardCellEmpty     uint8 = 0
	boardCellOwn             = 1
	boardCellFoe             = 2
	boardCellForbidden       = 3
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
			brd.Cells[i][j] = boardCellEmpty
		}
	}
	brd.Size = size
	return nil
}
