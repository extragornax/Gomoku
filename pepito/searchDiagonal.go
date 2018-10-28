package main

import (
	"gomoku/gomoku"
)

func checkDiagonalDown(game gomoku.Board, x uint, y uint, infos *searchResult, who uint8) {
	var count uint
	var maxSize uint
	j := x
	for i := y; i < game.Size && maxSize <= 5; i++ {
		if i < 0 || j < 0 || j >= (game.Size) {
			break
		}
		if game.Cells[i][j] == who {
			count++
		}
		j++
		maxSize++
	}
	if count > 0 {
		checkDiagonalDownEmpty(game, who, infos, x, y)
	} else {
		return
	}
	if infos.size < count && (infos.blockedBefore != true || infos.blockedAfter != true) {
		infos.x = x
		infos.y = y
		infos.size = count
	}
	return
}

func checkDiagonalDownEmpty(game gomoku.Board, who uint8, infos *searchResult, x uint, y uint) {
	if x+1 < game.Size && y+1 < game.Size &&
		game.Cells[y+1][x+1] != gomoku.BoardCellEmpty {
		infos.blockedAfter = true
	} else {
		infos.blockedAfter = false
	}
	if x-1 >= 0 && y-1 >= 0 &&
		game.Cells[y-1][x-1] != gomoku.BoardCellEmpty {
		infos.blockedBefore = true
	} else {
		infos.blockedBefore = false
	}
}

func searchDigonalDown(game gomoku.Board, who uint8) searchResult {
	var infos searchResult
	infos.resultType = searchDiagDownResult
	infos.x = 0
	infos.y = 0
	infos.size = 0
	for y := uint(0); y < game.Size; y++ {
		for x := uint(0); x < game.Size; x++ {
			if game.Cells[y][x] == who {
				checkDiagonalDown(game, x, y, &infos, who)
			}
		}
	}
	return infos
}

/*
**	Diagonal Up Functions
 */

func checkDiagonalUp(game gomoku.Board, x uint, y uint, infos *searchResult, who uint8) {
	var count uint
	var maxSize uint
	j := x
	for i := y; i < game.Size && maxSize <= 5; i-- {
		if i < 0 || j < 0 || j >= (game.Size) {
			break
		}
		if game.Cells[i][j] == who {
			count++
		} else {
			break
		}
		j++
		maxSize++
	}
	if count > 0 {
		checkDiagonalUpEmpty(game, who, infos, x, y)
	} else {
		return
	}
	if infos.size < count && (infos.blockedBefore != true || infos.blockedAfter != true) {
		infos.x = x
		infos.y = y
		infos.size = count
	}
	return
}

func checkDiagonalUpEmpty(game gomoku.Board, who uint8, infos *searchResult, x uint, y uint) {
	if (x+(infos.size+1)) < game.Size && (y-(infos.size+1)) >= 0 {
		if game.Cells[y-(infos.size+1)][x+(infos.size+1)] != gomoku.BoardCellEmpty &&
			game.Cells[y-(infos.size+1)][x+(infos.size+1)] != who {
			infos.blockedAfter = true
		} else {
			infos.blockedAfter = false
		}
	}
	if x == 0 {
		infos.blockedBefore = true
	} else if (x-1) >= 0 && (y+1) < game.Size {
		if game.Cells[y+1][x-1] != gomoku.BoardCellEmpty && game.Cells[y+1][x-1] != who {
			infos.blockedBefore = true
		} else {
			infos.blockedBefore = false
		}
	}
	if infos.blockedBefore == true || infos.blockedAfter == true {
		infos.size = 0
	}
}

func searchDigonalUp(game gomoku.Board, who uint8) searchResult {
	var infos searchResult
	infos.resultType = searchDiagUpResult
	infos.x = 0
	infos.y = 0
	infos.size = 0
	for y := uint(0); y < game.Size; y++ {
		for x := uint(0); x < game.Size; x++ {
			if game.Cells[y][x] == who {
				checkDiagonalUp(game, x, y, &infos, who)
			}
		}
	}
	return infos
}
