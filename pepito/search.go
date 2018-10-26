package main

import "gomoku/gomoku"

const (
	searchHorizontalResult = iota
	searchVecticalResult   = iota
	searchDiag1Result      = iota
	searchDiag2Result      = iota
)

// searchResult
type searchResult struct {
	resultType int
	size       uint
	x          uint
	y          uint
}

func searchHorizontal(brd *gomoku.Board, player uint8) searchResult {
	var res searchResult
	res.resultType = searchHorizontalResult
	res.size = 0
	res.x = 0
	res.y = 0
	piece := uint(0)

	for i := uint(0); i < brd.Size; i++ {
		for j := uint(0); j < brd.Size; j++ {
			for brd.Cells[i][j] == player {
				piece++
				j++
			}
			if piece > res.size {
				res.size = piece
				res.y = i
				res.x = j - piece
			}
			piece = 0
		}
	}
	return res
}

func searchVertical(brd *gomoku.Board, player uint8) searchResult {
	var res searchResult
	res.resultType = searchVecticalResult
	res.size = 0
	res.x = 0
	res.y = 0
	piece := uint(0)

	for j := uint(0); j < brd.Size; j++ {
		for i := uint(0); i < brd.Size; i++ {
			for brd.Cells[i][j] == player {
				piece++
				i++
			}
			if piece > res.size {
				res.size = piece
				res.y = i - piece
				res.x = j
			}
			piece = 0
		}
	}
	return res
}
