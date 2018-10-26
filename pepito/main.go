package main

import (
	"gomoku/gomoku"
)

type uvector struct {
	x uint
	y uint
}

func main() {
	var game gomoku.Gomoku
	game.Init()
	game.Run()
	for game.Live {
		trial := decisionTake(&game)
		game.Play(trial.x, trial.y)
		game.Run()
	}
}
