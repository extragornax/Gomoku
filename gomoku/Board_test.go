package gomoku

import "testing"

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
