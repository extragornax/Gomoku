package gomoku

import (
	"strings"
	"testing"
)

func TestMessageProcessorBoard(t *testing.T) {
	var gmk Gomoku

	gmk.Init()

	err := gmk.Board.init(5)
	if err != nil {
		t.Error("failed to init board")
	}

	messageProcessorBoardLine(&gmk, "1 1 2")
	if gmk.Board.Cells[1][1] != 2 {
		t.Error("board cell value is incorrect")
	}

	messageProcessorBoardLine(&gmk, "1 2 1")
	if gmk.Board.Cells[2][1] != 1 {
		t.Error("board cell value is incorrect")
	}

	messageProcessorBoardLine(&gmk, "4 0 2")
	if gmk.Board.Cells[0][4] != 2 {
		t.Error("board cell value is incorrect")
	}
}

func TestMessageProcessorTurn(t *testing.T) {
	var gmk Gomoku

	gmk.Init()

	err := gmk.Board.init(5)
	if err != nil {
		t.Error("failed to init board")
	}
	words := strings.Fields("TURN 3,2")
	messageProcessorTurn(&gmk, words[1:])
	if gmk.Board.Cells[2][3] != BoardCellFoe {
		t.Error("board cell value is incorrect")
	}
	if gmk.Turn != true {
		t.Error("turn should be true")
	}
	if gmk.Live != true {
		t.Error("gomoku should be up")
	}
}
