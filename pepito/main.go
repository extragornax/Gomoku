package main

import (
	"gomoku/gomoku"
)

func completeHorizontal(res searchResult) {
	if res.resultType == searchHorizontalResult {

	}
}

func main() {
	var game gomoku.Gomoku
	game.Init()
	game.Run()
	for game.Live {
		result := searchHorizontal(game.Board, gomoku.BoardCellFoe)
		tmp := searchVertical(game.Board, gomoku.BoardCellFoe)
		if tmp.size > result.size {
			result = tmp
		}

		game.Run()
	}
	game.Play(2, 2)
}
