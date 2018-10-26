package main

import (
	"gomoku/gomoku"
)

const (
	searchResultHorizontal = iota
	searchResultVectical   = iota
	searchDiagUpResult     = iota
	searchDiagDownResult   = iota
)

// searchResult
type searchResult struct {
	resultType    int
	size          uint
	x             uint
	y             uint
	blockedBefore bool
	blockedAfter  bool
}

func searchBlockedHorizontal(brd gomoku.Board, x uint, y uint, size uint) (before bool, after bool) {
	before = x == 0 || brd.Cells[y][x-1] != uint8(0)
	after = x+size+1 == brd.Size || brd.Cells[y][x+size] != uint8(0)
	return
	// return !(x == 0 || brd.Cells[y][x-1] != 0) ||
	// 	(x+size+1 == brd.Size || brd.Cells[y][x+size+1] != 0)
}

func searchBlockedVertical(brd gomoku.Board, x uint, y uint, size uint) (before bool, after bool) {
	before = y == 0 || brd.Cells[y-1][x] != 0
	after = y+size == brd.Size || brd.Cells[y+size][x] != 0
	return
}

func searchHorizontal(brd gomoku.Board, target uint8) searchResult {
	var res searchResult
	res.resultType = searchResultHorizontal
	res.size = 0
	res.x = 0
	res.y = 0

	for y := uint(0); y < brd.Size; y++ {
		for x := uint(0); x < brd.Size; x++ {
			size := uint(0)
			for brd.Cells[y][x] == target {
				size++
				x++
			}
			if size > res.size {
				before, after := searchBlockedHorizontal(brd, x-size, y, size)
				if !(before && after) {
					res.size = size
					res.y = y
					res.x = x - size
					res.blockedBefore = before
					res.blockedAfter = after
				}
			}
		}
	}
	return res
}

func searchVertical(brd gomoku.Board, player uint8) searchResult {
	var res searchResult
	res.resultType = searchResultVectical
	res.size = 0
	res.x = 0
	res.y = 0
	size := uint(0)

	for x := uint(0); x < brd.Size; x++ {
		for y := uint(0); y < brd.Size; y++ {
			for brd.Cells[y][x] == player {
				size++
				y++
			}
			if size > res.size {
				before, after := searchBlockedVertical(brd, x, y-size, size)
				if !(before && after) {
					res.size = size
					res.y = y - size
					res.x = x
					res.blockedBefore = before
					res.blockedAfter = after
				}
			}
			size = 0
		}
	}
	return res
}

func seachDiagonalUp(brd gomoku.Board, target uint8) searchResult {
	var res searchResult
	res.resultType = searchDiagUpResult

	return res
}

func seachDiagonalDown(brd gomoku.Board, target uint8) searchResult {
	var res searchResult
	res.resultType = searchDiagDownResult

	return res
}
