package main

import (
	"gomoku/gomoku"
	"testing"
)

func TestIter(t *testing.T) {
	var gmk gomoku.Gomoku

	gmk.Init()

	err := gmk.Board.Init(19)
	if err != nil {
		t.Error("Board not init")
	}

	total := toto(gmk.Board)

	if total != 1 {
		t.Error("Got return code ", total)
	}
}
