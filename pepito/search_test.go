package main

import (
	"gomoku/gomoku"
	"testing"
)

func TestSearchHorizontal(t *testing.T) {
	var board gomoku.Board
	var coord searchResult
	err := board.Init(10)
	if err != nil {
		t.Error("error was supposed to be nil")
	}
	// board.Cells[1][1] = gomoku.BoardCellFoe
	board.Cells[1][5] = gomoku.BoardCellFoe

	board.Cells[4][2] = gomoku.BoardCellFoe
	board.Cells[4][1] = gomoku.BoardCellFoe
	board.Cells[4][3] = gomoku.BoardCellOwn

	board.Cells[1][2] = gomoku.BoardCellOwn
	board.Cells[1][3] = gomoku.BoardCellOwn
	board.Cells[1][4] = gomoku.BoardCellOwn

	board.Cells[2][1] = gomoku.BoardCellOwn
	board.Cells[2][2] = gomoku.BoardCellOwn
	board.Cells[3][1] = gomoku.BoardCellOwn

	coord = searchHorizontal(board, gomoku.BoardCellOwn)
	if coord.size != 3 || coord.x != 2 || coord.y != 1 {
		t.Error("own combinasion could not be found")
	}
	if coord.blockedAfter != true || coord.blockedBefore != false {
		t.Error("blockade is incorrect")
	}
	coord = searchHorizontal(board, gomoku.BoardCellFoe)
	if coord.size != 2 || coord.x != 1 || coord.y != 4 {
		t.Error("foe's combinasion could not be found")
	}
	if coord.blockedAfter != true && coord.blockedBefore != false {
		t.Error("blockade is incorrect")
	}
}

func TestSearchVertical(t *testing.T) {
	var board gomoku.Board
	var coord searchResult
	err := board.Init(10)
	if err != nil {
		t.Error("error was supposed to be nil")
	}
	board.Cells[2][0] = gomoku.BoardCellOwn

	board.Cells[1][1] = gomoku.BoardCellFoe
	board.Cells[3][2] = gomoku.BoardCellFoe

	board.Cells[0][2] = gomoku.BoardCellFoe
	board.Cells[1][2] = gomoku.BoardCellOwn
	board.Cells[2][2] = gomoku.BoardCellOwn
	board.Cells[3][2] = gomoku.BoardCellOwn
	// board.Cells[4][2] = gomoku.BoardCellFoe

	board.Cells[2][3] = gomoku.BoardCellFoe
	board.Cells[3][3] = gomoku.BoardCellFoe

	coord = searchVertical(board, gomoku.BoardCellOwn)
	if coord.size != 3 || coord.x != 2 || coord.y != 1 {
		t.Error("error coord hor")
	}
	if coord.blockedAfter != false && coord.blockedBefore != true {
		t.Error("blockade is incorrect")
	}
	coord = searchVertical(board, gomoku.BoardCellFoe)
	if coord.size != 2 || coord.x != 3 || coord.y != 2 {
		t.Error("error coord hor")
	}
	if coord.blockedAfter != false && coord.blockedBefore != false {
		t.Error("blockade is incorrect")
	}
}

func TestSearchVerticalBlocked(t *testing.T) {
	var board gomoku.Board
	var coord searchResult
	err := board.Init(10)
	if err != nil {
		t.Error("error was supposed to be nil")
	}
	board.Cells[2][0] = gomoku.BoardCellOwn

	board.Cells[1][1] = gomoku.BoardCellFoe
	board.Cells[3][2] = gomoku.BoardCellFoe

	board.Cells[0][2] = gomoku.BoardCellFoe
	board.Cells[1][2] = gomoku.BoardCellOwn
	board.Cells[2][2] = gomoku.BoardCellOwn
	board.Cells[3][2] = gomoku.BoardCellOwn
	board.Cells[4][2] = gomoku.BoardCellFoe

	board.Cells[2][3] = gomoku.BoardCellFoe
	board.Cells[3][3] = gomoku.BoardCellFoe

	coord = searchVertical(board, gomoku.BoardCellOwn)
	if coord.size != 1 || coord.x != 0 || coord.y != 2 {
		t.Error("error coord hor")
	}
	coord = searchVertical(board, gomoku.BoardCellFoe)
	if coord.size != 2 || coord.x != 3 || coord.y != 2 {
		t.Error("error coord hor")
	}
}

func TestSearchHorizontalSingle(t *testing.T) {
	var board gomoku.Board
	var coord searchResult
	err := board.Init(10)
	if err != nil {
		t.Error("error was supposed to be nil")
	}
	board.Cells[0][5] = gomoku.BoardCellOwn
	board.Cells[1][1] = gomoku.BoardCellOwn
	board.Cells[1][4] = gomoku.BoardCellFoe
	board.Cells[2][2] = gomoku.BoardCellFoe
	board.Cells[3][1] = gomoku.BoardCellFoe
	board.Cells[3][3] = gomoku.BoardCellOwn

	coord = searchHorizontal(board, gomoku.BoardCellOwn)
	if coord.size != 1 || coord.x != 5 || coord.y != 0 {
		t.Error("error coord hor")
	}
	coord = searchHorizontal(board, gomoku.BoardCellFoe)
	if coord.size != 1 || coord.x != 4 || coord.y != 1 {
		t.Error("error coord hor")
	}
}

func TestSearchHorizontalZero(t *testing.T) {
	var board gomoku.Board
	var coord searchResult
	err := board.Init(10)
	if err != nil {
		t.Error("error was supposed to be nil")
	}

	coord = searchHorizontal(board, gomoku.BoardCellOwn)
	if coord.size != 0 {
		t.Error("error coord hor")
	}
	coord = searchHorizontal(board, gomoku.BoardCellFoe)
	if coord.size != 0 {
		t.Error("error coord hor")
	}
}

func TestSearchVerticalSingle(t *testing.T) {
	var board gomoku.Board
	var coord searchResult
	err := board.Init(10)
	if err != nil {
		t.Error("error was supposed to be nil")
	}
	board.Cells[0][5] = gomoku.BoardCellOwn
	board.Cells[1][1] = gomoku.BoardCellOwn
	board.Cells[1][4] = gomoku.BoardCellFoe
	board.Cells[2][2] = gomoku.BoardCellFoe
	board.Cells[3][1] = gomoku.BoardCellFoe
	board.Cells[3][3] = gomoku.BoardCellOwn

	coord = searchVertical(board, gomoku.BoardCellOwn)
	if coord.size != 1 || coord.x != 1 || coord.y != 1 {
		t.Error("error coord hor")
	}
	coord = searchVertical(board, gomoku.BoardCellFoe)
	if coord.size != 1 || coord.x != 1 || coord.y != 3 {
		t.Error("error coord hor")
	}
}

func TestSearchVerticalZero(t *testing.T) {
	var board gomoku.Board
	var coord searchResult
	err := board.Init(10)
	if err != nil {
		t.Error("error was supposed to be nil")
	}

	coord = searchVertical(board, gomoku.BoardCellOwn)
	if coord.size != 0 {
		t.Error("error coord hor")
	}
	coord = searchVertical(board, gomoku.BoardCellFoe)
	if coord.size != 0 {
		t.Error("error coord hor")
	}
}

func TestSearchDiagDownBlockedAfter(t *testing.T) {
	var board gomoku.Board
	var coord searchResult
	err := board.Init(10)
	if err != nil {
		t.Error("error was supposed to be nil")
	}

	// board.Cells[0][1] = gomoku.BoardCellFoe
	board.Cells[1][2] = gomoku.BoardCellOwn
	board.Cells[2][3] = gomoku.BoardCellOwn
	board.Cells[3][4] = gomoku.BoardCellOwn
	board.Cells[4][5] = gomoku.BoardCellFoe
	// board.Cells[3][1] = gomoku.BoardCellOwn
	// board.Cells[4][2] = gomoku.BoardCellOwn
	// board.Cells[5][3] = gomoku.BoardCellOwn

	coord = searchDiagonalDown(board, gomoku.BoardCellOwn)
	if coord.size != 3 || coord.x != 2 || coord.y != 1 {
		t.Error("error coord hor")
	}
	if !coord.blockedAfter || coord.blockedBefore {
		t.Error("bad blockade")
	}
}

func TestSearchDiagDownBlockedBefore(t *testing.T) {
	var board gomoku.Board
	var coord searchResult
	err := board.Init(10)
	if err != nil {
		t.Error("error was supposed to be nil")
	}

	board.Cells[0][1] = gomoku.BoardCellFoe
	board.Cells[1][2] = gomoku.BoardCellOwn
	board.Cells[2][3] = gomoku.BoardCellOwn
	board.Cells[3][4] = gomoku.BoardCellOwn
	// board.Cells[4][5] = gomoku.BoardCellFoe
	// board.Cells[3][1] = gomoku.BoardCellOwn
	// board.Cells[4][2] = gomoku.BoardCellOwn
	// board.Cells[5][3] = gomoku.BoardCellOwn

	coord = searchDiagonalDown(board, gomoku.BoardCellOwn)
	if coord.size != 3 || coord.x != 2 || coord.y != 1 {
		t.Error("error coord hor")
	}
	if coord.blockedAfter || !coord.blockedBefore {
		t.Error("bad blockade")
	}
}
