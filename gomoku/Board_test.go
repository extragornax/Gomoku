package gomoku

import (
	"testing"
)

func TestBoardCreateOK(t *testing.T) {
	var board Board

	err := board.init(20)
	if err != nil {
		t.Error("error was supposed to be nil")
	}
	if len(board.Cells) != 20 {
		t.Errorf("len was suppoesed to be 20 and it was %d", len(board.Cells))
	}
	if board.Size != 20 {
		t.Errorf("size was suppoesed to be 20 and it was %d", board.Size)
	}

	err = board.init(10)
	if err != nil {
		t.Error("error was supposed to be nil")
	}
	if len(board.Cells) != 10 {
		t.Errorf("len was suppoesed to be 10 and it was %d", len(board.Cells))
	}
	if board.Size != 10 {
		t.Errorf("size was suppoesed to be 10 and it was %d", board.Size)
	}
}

func TestBoardCreateKO(t *testing.T) {
	var board Board
	err := board.init(30)
	if err == nil || err.Error() != ErrBoardTooLarge {
		t.Error("Expected err board too large")
	}

	err = board.init(0)
	if err == nil || err.Error() != ErrBoardTooSmall {
		t.Error("Expected err board too small")
	}
}

func TestSearch_hor(t *testing.T) {
	var board Board
	var coord Coord
	err := board.init(10)
	if err != nil {
		t.Error("error was supposed to be nil")
	}
	board.Cells[1][1] = BoardCellFoe
	board.Cells[1][0] = BoardCellFoe
	board.Cells[1][2] = BoardCellOwn
	board.Cells[1][4] = BoardCellOwn
	board.Cells[1][3] = BoardCellOwn
	board.Cells[2][1] = BoardCellOwn
	board.Cells[2][2] = BoardCellOwn
	board.Cells[3][1] = BoardCellOwn

	coord = board.searchHor(BoardCellOwn)
	if coord.NbPiece != 3 || coord.CoordX != 2 || coord.CoordY != 1 {
		t.Error("error coord hor")
	}
	coord = board.searchHor(BoardCellFoe)
	if coord.NbPiece != 2 || coord.CoordX != 0 || coord.CoordY != 1 {
		t.Error("error coord hor")
	}
}

func TestSearch_ver(t *testing.T) {
	var board Board
	var coord Coord
	err := board.init(10)
	if err != nil {
		t.Error("error was supposed to be nil")
	}
	board.Cells[2][0] = BoardCellOwn
	board.Cells[1][1] = BoardCellFoe
	board.Cells[3][2] = BoardCellFoe
	board.Cells[1][2] = BoardCellOwn
	board.Cells[2][2] = BoardCellOwn
	board.Cells[3][2] = BoardCellOwn
	board.Cells[2][3] = BoardCellFoe
	board.Cells[3][3] = BoardCellFoe

	coord = board.searchVer(BoardCellOwn)
	if coord.NbPiece != 3 || coord.CoordX != 2 || coord.CoordY != 1 {
		t.Error("error coord hor")
	}
	coord = board.searchVer(BoardCellFoe)
	if coord.NbPiece != 2 || coord.CoordX != 3 || coord.CoordY != 2 {
		t.Error("error coord hor")
	}
}

func TestSearch_hor_single(t *testing.T) {
	var board Board
	var coord Coord
	err := board.init(10)
	if err != nil {
		t.Error("error was supposed to be nil")
	}
	board.Cells[0][5] = BoardCellOwn
	board.Cells[1][1] = BoardCellOwn
	board.Cells[1][4] = BoardCellFoe
	board.Cells[2][2] = BoardCellFoe
	board.Cells[3][1] = BoardCellFoe
	board.Cells[3][3] = BoardCellOwn

	coord = board.searchHor(BoardCellOwn)
	if coord.NbPiece != 1 || coord.CoordX != 5 || coord.CoordY != 0 {
		t.Error("error coord hor")
	}
	coord = board.searchHor(BoardCellFoe)
	if coord.NbPiece != 1 || coord.CoordX != 4 || coord.CoordY != 1 {
		t.Error("error coord hor")
	}
}

func TestSearch_ver_single(t *testing.T) {
	var board Board
	var coord Coord
	err := board.init(10)
	if err != nil {
		t.Error("error was supposed to be nil")
	}
	board.Cells[0][5] = BoardCellOwn
	board.Cells[1][1] = BoardCellOwn
	board.Cells[1][4] = BoardCellFoe
	board.Cells[2][2] = BoardCellFoe
	board.Cells[3][1] = BoardCellFoe
	board.Cells[3][3] = BoardCellOwn

	coord = board.searchVer(BoardCellOwn)
	if coord.NbPiece != 1 || coord.CoordX != 1 || coord.CoordY != 1 {
		t.Error("error coord hor")
	}
	coord = board.searchVer(BoardCellFoe)
	if coord.NbPiece != 1 || coord.CoordX != 1 || coord.CoordY != 3 {
		t.Error("error coord hor")
	}
}
