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

// func check_diagonal_up(game gomoku.Gomoku, x uint, y uint, ex_count uint, who uint8) (int, int, uint) {
// 	var count uint = 0
// 	for i := y; i < game.Board.Size; i++ {
// 		for j := x; j < game.Board.Size; j-- {
// 			if i < 0 || j < 0 || i >= (game.Board.Size) {
// 				return int(0), int(0), ex_count
// 			}
// 			if game.Board.Cells[i][j] == who {
// 				count += 1
// 			}
// 			if count > ex_count+1 && count < 5 {
// 				return int(x), int(y), count
// 			}
// 		}
// 	}
// 	return int(-1), int(-1), ex_count
// }

// func find_diag_up(game gomoku.Gomoku, who uint8) {
// 	var max_y int = -1
// 	var max_x int = -1
// 	var max_count uint = 0
// 	for i := uint(0); i < game.Board.Size; i++ {
// 		for j := uint(0); j < game.Board.Size; j++ {
// 			if game.Board.Cells[i][j] == who {
// 				tmp_x, tmp_y, tmp_count := check_diagonal_up(game, j, i, max_count, who)
// 				if tmp_count > max_count {
// 					max_count = tmp_count
// 					max_x = tmp_x
// 					max_y = tmp_y
// 					fmt.Printf("DIAGONAL UP: %d %d %d\n", max_x, max_y, max_count)
// 				}
// 			}
// 		}
// 	}
// }

func check_diagonal_up(game gomoku.Gomoku, x uint, y uint, ex_count uint, who uint8) uint {
	var count uint = 0
	var j uint = x
	for i := int8(y); i < int8(game.Board.Size); i-- {
		if i < 0 || j < 0 || j >= (game.Board.Size) {
			return ex_count
		}
		if game.Board.Cells[i][j] == who {
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
				tmp_count := check_diagonal_up(game, j, i, max_count, who)
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

func check_diagonal_down(game gomoku.Gomoku, x uint, y uint, ex_count uint, who uint8) uint {
	var count uint = 0
	var j uint = x
	for i := int8(y); i < int8(game.Board.Size); i++ {
		if i < 0 || j < 0 || j >= (game.Board.Size) {
			return ex_count
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
				tmp_count := check_diagonal_down(game, j, i, max_count, who)
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
