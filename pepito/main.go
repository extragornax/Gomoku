package main

import (
	"fmt"
	"gomoku/gomoku"
	"math/rand"
)

func kappa(gmk gomoku.Gomoku, r *rand.Rand) {
	var a uint = uint(r.Intn(19))
	var b uint = uint(r.Intn(19))
	gmk.Play(a, b)
}
func check_diagonal_up(game gomoku.Gomoku, x uint, y uint, infos *searchResult, who uint8) uint {
	var count uint = 0
	var j uint = x
	for i := int8(y); i < int8(game.Board.Size); i-- {
		if i < 0 || j < 0 || j >= (game.Board.Size) {
			return count
		}
		if game.Board.Cells[i][j] == who {
			count += 1
		}
		j += 1
	}
	if count > infos.size {
		infos.size = count
		infos.x = x
		infos.y = y
	}
	return count
}

func check_empty_before_after_up(game gomoku.Gomoku, who uint8, infos *searchResult) {
	if infos.x+1 < game.Board.Size && infos.y >= 0 &&
		game.Board.Cells[infos.x+1][infos.y-1] == who {
		infos.blockedAfter = false
	} else {
		infos.blockedAfter = true
	}
	if infos.x-1 >= 0 && infos.y+1 < game.Board.Size &&
		game.Board.Cells[infos.x-1][infos.y+1] == who {
		infos.blockedBefore = false
	} else {
		infos.blockedBefore = true
	}
}

func find_diag_up(game gomoku.Gomoku, who uint8) searchResult {
	var infos searchResult
	infos.resultType = searchDiagUpResult
	infos.x = 0
	infos.y = 0
	for i := uint(0); i < game.Board.Size; i++ {
		for j := uint(0); j < game.Board.Size; j++ {
			if game.Board.Cells[i][j] == who {
				check_diagonal_up(game, j, i, &infos, who)
			}
		}
	}
	if infos.size > 0 {
		check_empty_before_after_up(game, who, &infos)
	}
	// fmt.Printf("DIAGONAL UP: %d %d %d %d\n", infos.x, infos.y, infos.size, infos.blockedBefore)
	return infos
}

type uvector struct {
	x uint
	y uint
}

func main() {
	var game gomoku.Gomoku
	game.Init()
	// r := rand.New(rand.NewSource(99))
	game.Run()
	var i uint = 6
	var j uint = 6
	for game.Live {
		j += 1
		i -= 1
		game.Play(i, j)
		// find_diag_up(game, gomoku.BoardCellFoe)
		fmt.Printf("==========================\n")
		for i := uint(0); i < game.Board.Size; i++ {
			for h := uint(0); h < game.Board.Size; h++ {
				fmt.Printf("%d", game.Board.Cells[i][h])
			}
			fmt.Printf("\n")
		}
		fmt.Printf("==========================\n")
		searchDigonalDown(game.Board, gomoku.BoardCellFoe)
		game.Run()
	}
}
