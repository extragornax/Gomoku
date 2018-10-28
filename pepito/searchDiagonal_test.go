package main

import (
	"gomoku/gomoku"
	"testing"
)

func TestSearchDiagonalDownSingle(t *testing.T) {
	var board gomoku.Board
	var coord searchResult
	err := board.Init(10)
	if err != nil {
		t.Error("error was supposed to be nil")
	}
	board.Cells[5][5] = gomoku.BoardCellOwn
	board.Cells[4][4] = gomoku.BoardCellFoe

	coord = searchDigonalDown(board, gomoku.BoardCellFoe)
	if coord.size != 1 || coord.x != 4 || coord.y != 4 || coord.blockedAfter != true || coord.blockedBefore != false {
		t.Error("error coord hor")
	}
	coord = searchDigonalDown(board, gomoku.BoardCellOwn)
	if coord.size != 1 || coord.x != 5 || coord.y != 5 || coord.blockedAfter != false || coord.blockedBefore != true {
		t.Error("error coord hor")
	}
}

func TestSearchDiagonalDownFour(t *testing.T) {
	var board gomoku.Board
	var coord searchResult
	err := board.Init(10)
	if err != nil {
		t.Error("error was supposed to be nil")
	}
	board.Cells[5][5] = gomoku.BoardCellOwn
	board.Cells[4][4] = gomoku.BoardCellFoe
	board.Cells[3][3] = gomoku.BoardCellFoe
	board.Cells[2][2] = gomoku.BoardCellFoe
	board.Cells[1][1] = gomoku.BoardCellFoe

	coord = searchDigonalDown(board, gomoku.BoardCellFoe)
	if coord.size != 4 || coord.x != 1 || coord.y != 1 {
		t.Error("error coord hor")
	}
}

func TestSearchDiagonalDownFourBlocked(t *testing.T) {
	var board gomoku.Board
	var coord searchResult
	err := board.Init(10)
	if err != nil {
		t.Error("error was supposed to be nil")
	}
	board.Cells[0][0] = gomoku.BoardCellOwn
	board.Cells[5][5] = gomoku.BoardCellOwn
	board.Cells[4][4] = gomoku.BoardCellFoe
	board.Cells[3][3] = gomoku.BoardCellFoe
	board.Cells[2][2] = gomoku.BoardCellFoe
	board.Cells[1][1] = gomoku.BoardCellFoe

	coord = searchDigonalDown(board, gomoku.BoardCellFoe)
	if coord.size != 0 {
		t.Error("error coord hor")
	}
}

func TestSearchDiagonalDownFourBlockedOneFree(t *testing.T) {
	var board gomoku.Board
	var coord searchResult
	err := board.Init(10)
	if err != nil {
		t.Error("error was supposed to be nil")
	}
	board.Cells[0][0] = gomoku.BoardCellOwn
	board.Cells[5][5] = gomoku.BoardCellOwn
	board.Cells[4][4] = gomoku.BoardCellFoe
	board.Cells[3][3] = gomoku.BoardCellFoe
	board.Cells[2][2] = gomoku.BoardCellFoe
	board.Cells[1][1] = gomoku.BoardCellFoe
	board.Cells[8][8] = gomoku.BoardCellFoe

	coord = searchDigonalDown(board, gomoku.BoardCellFoe)
	if coord.size != 1 || coord.x != 8 || coord.y != 8 {
		t.Error("error coord hor")
	}
}

/*
** Test Up
 */
func TestSearchDiagonalUpSingle(t *testing.T) {
	var board gomoku.Board
	var coord searchResult
	err := board.Init(10)
	if err != nil {
		t.Error("error was supposed to be nil")
	}
	board.Cells[5][5] = gomoku.BoardCellOwn
	board.Cells[6][4] = gomoku.BoardCellFoe

	coord = searchDigonalUp(board, gomoku.BoardCellFoe)
	if coord.size != 1 || coord.x != 4 || coord.y != 6 || coord.blockedAfter != true || coord.blockedBefore != false {
		t.Error("error coord hor foe")
	}
	coord = searchDigonalUp(board, gomoku.BoardCellOwn)
	if coord.size != 1 || coord.x != 5 || coord.y != 5 || coord.blockedAfter != false || coord.blockedBefore != true {
		t.Error("error coord hor")
	}
}

func TestSearchDiagonalUpFour(t *testing.T) {
	var board gomoku.Board
	var coord searchResult
	err := board.Init(10)
	if err != nil {
		t.Error("error was supposed to be nil")
	}
	board.Cells[5][5] = gomoku.BoardCellOwn
	board.Cells[4][4] = gomoku.BoardCellFoe
	board.Cells[3][3] = gomoku.BoardCellFoe
	board.Cells[2][2] = gomoku.BoardCellFoe
	board.Cells[1][1] = gomoku.BoardCellFoe

	coord = searchDigonalUp(board, gomoku.BoardCellFoe)
	if coord.size != 4 || coord.x != 1 || coord.y != 1 {
		t.Error("error coord hor")
	}
}

func TestSearchDiagonalUpFourBlocked(t *testing.T) {
	var board gomoku.Board
	var coord searchResult
	err := board.Init(10)
	if err != nil {
		t.Error("error was supposed to be nil")
	}
	board.Cells[0][0] = gomoku.BoardCellOwn
	board.Cells[5][5] = gomoku.BoardCellOwn
	board.Cells[4][4] = gomoku.BoardCellFoe
	board.Cells[3][3] = gomoku.BoardCellFoe
	board.Cells[2][2] = gomoku.BoardCellFoe
	board.Cells[1][1] = gomoku.BoardCellFoe

	coord = searchDigonalUp(board, gomoku.BoardCellFoe)
	if coord.size != 0 {
		t.Error("error coord hor")
	}
}

func TestSearchDiagonalUpFourBlockedOneFree(t *testing.T) {
	var board gomoku.Board
	var coord searchResult
	err := board.Init(10)
	if err != nil {
		t.Error("error was supposed to be nil")
	}
	board.Cells[0][0] = gomoku.BoardCellOwn
	board.Cells[5][5] = gomoku.BoardCellOwn
	board.Cells[4][4] = gomoku.BoardCellFoe
	board.Cells[3][3] = gomoku.BoardCellFoe
	board.Cells[2][2] = gomoku.BoardCellFoe
	board.Cells[1][1] = gomoku.BoardCellFoe
	board.Cells[8][8] = gomoku.BoardCellFoe

	coord = searchDigonalUp(board, gomoku.BoardCellFoe)
	if coord.size != 1 || coord.x != 8 || coord.y != 8 {
		t.Error("error coord hor")
	}
}
