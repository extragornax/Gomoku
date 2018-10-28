package main

import (
	"fmt"
	"gomoku/gomoku"
)

func checkDiagonalDown(game gomoku.Gomoku, x uint, y uint, infos *searchResult, who uint8) uint {
	var count uint
	var maxSize uint
	j := x
	for i := y; i < game.Board.Size && maxSize <= 5; i++ {
		if i < 0 || j < 0 || j >= (game.Board.Size) {
			if infos.size < count {
				infos.x = x
				infos.y = y
				infos.size = count
			}
			return count
		}
		if game.Board.Cells[i][j] == who {
			count++
		}
		j++
		maxSize++
	}
	if infos.size < count {
		infos.x = x
		infos.y = y
		infos.size = count
	}
	return count
}

func checkDiagonalDownEmpty(game gomoku.Gomoku, who uint8, infos *searchResult) {
	if infos.x+1 < game.Board.Size && infos.y+1 < game.Board.Size &&
		game.Board.Cells[infos.y+1][infos.x+1] != gomoku.BoardCellEmpty {
		infos.blockedAfter = true
	} else {
		infos.blockedAfter = false
	}
	if infos.x-1 >= 0 && infos.y-1 >= 0 &&
		game.Board.Cells[infos.y-1][infos.x-1] != gomoku.BoardCellEmpty {
		infos.blockedBefore = true
	} else {
		infos.blockedBefore = false
	}
}

func searchDigonalDown(game gomoku.Gomoku, who uint8) searchResult {
	var infos searchResult
	infos.resultType = searchDiagDownResult
	infos.x = 0
	infos.y = 0
	infos.size = 0
	for y := uint(0); y < game.Board.Size; y++ {
		for x := uint(0); x < game.Board.Size; x++ {
			if game.Board.Cells[y][x] == who {
				checkDiagonalDown(game, x, y, &infos, who)
			}
		}
	}
	if infos.size > 0 {
		checkDiagonalDownEmpty(game, who, &infos)
	}
	fmt.Printf("DIAGONAL DOWN: %d %d %d before:%d after:%d\n", infos.x, infos.y, infos.size, infos.blockedBefore, infos.blockedAfter)
	return infos
}
