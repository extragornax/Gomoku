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
func (brd *Board) Init(size uint) error {
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

type Coord struct {
	NbPiece uint
	CoordX  uint
	CoordY  uint
}

func (brd *Board) searchHor(player uint8) Coord {
	var coor Coord
	coor.NbPiece = 0
	coor.CoordX = 0
	coor.CoordY = 0
	piece := uint(0)

	for i := uint(0); i < brd.Size; i++ {
		for j := uint(0); j < brd.Size; j++ {
			for brd.Cells[i][j] == player {
				piece++
				j++
			}
			if piece > coor.NbPiece {
				coor.NbPiece = piece
				coor.CoordY = i
				coor.CoordX = j - piece
			}
			piece = 0
		}
	}
	return coor
}

func (brd *Board) searchVer(player uint8) Coord {
	var coor Coord
	coor.NbPiece = 0
	coor.CoordX = 0
	coor.CoordY = 0
	piece := uint(0)

	for j := uint(0); j < brd.Size; j++ {
		for i := uint(0); i < brd.Size; i++ {
			for brd.Cells[i][j] == player {
				piece++
				i++
			}
			if piece > coor.NbPiece {
				coor.NbPiece = piece
				coor.CoordY = i - piece
				coor.CoordX = j
			}
			piece = 0
		}
	}
	return coor
}
