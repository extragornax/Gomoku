package main

import (
	"gomoku/gomoku"
	"testing"
)

func TestDecisionCompleOrBlock1(t *testing.T) {
	var game gomoku.Gomoku
	game.Init()
	err := game.Board.Init(10)
	if err != nil {
		t.Error("error was supposed to be nil")
	}

	game.Board.Cells[3][3] = gomoku.BoardCellFoe
	game.Board.Cells[3][4] = gomoku.BoardCellFoe
	game.Board.Cells[3][5] = gomoku.BoardCellFoe
	game.Board.Cells[3][1] = gomoku.BoardCellOwn
	game.Board.Cells[4][1] = gomoku.BoardCellOwn
	game.Board.Cells[5][1] = gomoku.BoardCellOwn
	decision := decisionTake(&game)
	if decision.x != 2 && decision.y != 3 {
		t.Errorf("Not on the expected tile, got: %d, %d", decision.x, decision.y)
	}
}

func TestDecisionCompleOrBlock2(t *testing.T) {
	var game gomoku.Gomoku
	game.Init()
	err := game.Board.Init(10)
	if err != nil {
		t.Error("error was supposed to be nil")
	}

	game.Board.Cells[3][3] = gomoku.BoardCellFoe
	game.Board.Cells[3][4] = gomoku.BoardCellFoe
	game.Board.Cells[3][5] = gomoku.BoardCellOwn
	game.Board.Cells[3][1] = gomoku.BoardCellOwn
	game.Board.Cells[4][1] = gomoku.BoardCellOwn
	game.Board.Cells[5][1] = gomoku.BoardCellOwn
	decision := decisionTake(&game)
	if decision.x != 1 && decision.y != 2 {
		t.Errorf("Not on the expected tile, got: %d, %d", decision.x, decision.y)
	}
}

func TestDecisionCompleOrBlock3(t *testing.T) {
	var game gomoku.Gomoku
	game.Init()
	err := game.Board.Init(10)
	if err != nil {
		t.Error("error was supposed to be nil")
	}

	game.Board.Cells[3][3] = gomoku.BoardCellFoe
	game.Board.Cells[3][4] = gomoku.BoardCellFoe
	game.Board.Cells[3][5] = gomoku.BoardCellOwn
	game.Board.Cells[2][1] = gomoku.BoardCellFoe
	game.Board.Cells[3][1] = gomoku.BoardCellOwn
	game.Board.Cells[4][1] = gomoku.BoardCellOwn
	game.Board.Cells[5][1] = gomoku.BoardCellOwn
	decision := decisionTake(&game)
	if decision.x != 1 && decision.y != 6 {
		t.Errorf("Not on the expected tile, got: %d, %d", decision.x, decision.y)
	}
}
