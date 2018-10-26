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
func check_diagonal_up(game gomoku.Gomoku, x uint, y uint, who uint8) uint {
	var count uint = 0
	var j uint = x
	for i := int8(y); i < int8(game.Board.Size); i-- {
		if i < 0 || j < 0 || j >= (game.Board.Size) {
			return count
		}
		if game.Board.Cells[i][j] == who {
			fmt.Printf("coucou = %d\n", count)
			count += 1
		}
		j += 1
	}
	return count
}

func find_diag_up(game gomoku.Gomoku, who uint8) {
	var max_y int = -1
	var max_x int = -1
	var max_count uint = 0
	for i := uint(0); i < game.Board.Size; i++ {
		for j := uint(0); j < game.Board.Size; j++ {
			if game.Board.Cells[i][j] == who {
				tmp_count := check_diagonal_up(game, j, i, who)
				fmt.Printf("%d\n", tmp_count)
				if tmp_count > max_count {
					max_count = tmp_count
					max_x = int(j)
					max_y = int(i)
					fmt.Printf("DIAGONAL UP: %d %d %d\n", max_x, max_y, max_count)
				}
			}
		}
	}
}

func check_diagonal_down(game gomoku.Gomoku, x uint, y uint, who uint8) uint {
	var count uint = 0
	var j uint = x
	for i := int8(y); i < int8(game.Board.Size); i++ {
		if i < 0 || j < 0 || j >= (game.Board.Size) {
			return count
		}
		if game.Board.Cells[i][j] == who {
			count += 1
		}
		j += 1
	}
	return count
}

func find_diag_down(game gomoku.Gomoku, who uint8) {
	var max_y int = -1
	var max_x int = -1
	var max_count uint = 0
	for i := uint(0); i < game.Board.Size; i++ {
		for j := uint(0); j < game.Board.Size; j++ {
			if game.Board.Cells[i][j] == who {
				tmp_count := check_diagonal_down(game, j, i, who)
				if tmp_count > max_count {
					max_count = tmp_count
					max_x = int(j)
					max_y = int(i)
					fmt.Printf("DIAGONAL DOWN: %d %d %d\n", max_x, max_y, max_count)
				}
			}
		}
	}
}

func main() {
	var game gomoku.Gomoku
	game.Init()
	r := rand.New(rand.NewSource(99))
	game.Run()
	for game.Live {
		kappa(game, r)
		find_diag_up(game, gomoku.BoardCellFoe)
		find_diag_down(game, gomoku.BoardCellFoe)
		game.Run()
	}
}
