package main

import (
	"fmt"
	"gomoku/gomoku"
)

const (
	// BoardCellEmpty is a cell where no one has placed anything
	BoardCellEmpty uint8 = 0
	// BoardCellOwn : your cell
	BoardCellOwn = 1
	// BoardCellFoe : your oponent's cell
	BoardCellFoe = 2
	// BoardCellForbidden : a cell where no one can play
	BoardCellForbidden = 3
	// My turn
	TurnOwn = true
	// Other player turn
	TurnFoe = false
)

func retCellValue(turn bool) uint8 {
	if turn == TurnOwn {
		return BoardCellOwn
	} else {
		return BoardCellFoe
	}
}

func retNextTurn(turn bool) bool {
	if turn == TurnOwn {
		return TurnFoe
	} else {
		return TurnOwn
	}
}

func printTab(cells [][]uint8, size uint, instance uint) {
	println("======================================")
	println(instance)
	for i := uint(0); i < size; i++ {
		for j := uint(0); j < size; j++ {
			print(cells[i][j])
		}
		print("\n")
	}
	println("======================================")
}

func copyCells(src [][]uint8) (cells [][]uint8) {
	size := uint(len(src))
	cells = make([][]uint8, size)
	for i := uint(0); i < size; i++ {
		cells[i] = append(cells[i], src[i]...)
		// for j := uint(0); j < size; j++ {
		// 	cells[i][j] =
		// }
	}
	return
}

func iter(cells [][]uint8, size uint, myturn bool, instance uint) uint {
	if instance == 2 {
		return uint(0)
	}
	i := uint(0)
	j := uint(0)
	for i < size {
		for j < size {
			if cells[i][j] == BoardCellEmpty {
				src := copyCells(cells)
				src[i][j] = retCellValue(myturn)
				//	fmt.Println("Rec ", i, " ", j)
				//printTab(src, size, instance)
				fmt.Print("e")
				iter(src, size, retNextTurn(myturn), instance+1)
			}
			j++
		}
		j = 0
		i++
	}
	return uint(0)
}

func toto(board gomoku.Board) uint {
	fmt.Print("Starting toto\n")
	for i := uint(0); i < board.Size; i++ {
		for j := uint(0); j < board.Size; j++ {
			if board.Cells[i][j] != BoardCellEmpty {
				board.Cells[i][j] = BoardCellEmpty
			}
		}
	}
	tata := iter(board.Cells, board.Size, TurnOwn, 0)
	fmt.Print("End toto\n")
	return tata
}
