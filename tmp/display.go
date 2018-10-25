package main

import (
	"fmt"
	"regexp"
	"time"
)

func displayBoard(tools *Tools, game *Game) {

	// putImageCenter("ressources/board.bmp", tools.surface)
	i := 0
	j := 0

	for i < SIZE {
		j = 0
		for j < SIZE {
			// if game.board[i][j] != 0 {
			// displayDot(int32(i), int32(j), tools, game.board[i][j])
			if game.board[i][j] != NONE {
				fmt.Print(game.board[i][j])
			} else {
				fmt.Print(" ")
			}
			// }
			j++
		}
		println()
		i++
	}
	i = 0
	for i < game.score[PLAYER1-1] {
		displayScoreDot(int32(i), tools, PLAYER2)
		i++
	}
	i = 0
	for i < game.score[PLAYER2-1] {
		displayScoreDot(int32(i), tools, PLAYER1)
		i++
	}
	displayTurn(tools, game.curPlayer)

}

func displayTurn(tools *Tools, curPlayer int) {
	var file string

	if curPlayer == PLAYER1 {
		file = "left"
	} else {
		file = "right"
	}
	fmt.Println("Turn ", file)
}

func displayDot(i int32, j int32, tools *Tools, player int) {
	var file string

	if player == PLAYER1 {
		file = "PLAYER1"
	} else if player == PLAYER2 {
		file = "PLAYER2"
	} else if player == HINT {
		file = "HINT"
	}

	fmt.Println("displaydot ", file)
}

func displayScoreDot(nb int32, tools *Tools, player int) {
	var file string

	if player == PLAYER1 {
		file = "PLAYER1"
	} else {
		file = "PLAYER2"
	}
	if nb == 9 {

		fmt.Println("nb 9 ", file)
	} else if player == PLAYER1 {
		fmt.Println("PLAYER 1 ", file, " nb ", nb)
	} else {
		fmt.Println("PLAYER ELSE ", file, " nb ", nb)
	}
}

func displayDepth(tools *Tools, sum int) {

	if sum < 0 && DEPTH_MAX+sum == 0 {
		sum = 0
	}
	if sum > 0 && DEPTH_MAX+sum == 11 {
		sum = 0
	}
	DEPTH_MAX += sum

	fmt.Println("depth ", DEPTH_MAX)
}

func displayTime(tools *Tools, time time.Duration) {

	match, _ := regexp.MatchString("µs", time.String())
	var str string

	if match {
		str = "less than 1ms"
	} else {
		str = time.String()
	}
	fmt.Println(str)
}

func endGame(tools *Tools, game *Game) {

	msg := ""
	if game.winner == PLAYER1 {
		msg = "PLAYER 1 win"
	} else if game.winner == PLAYER2 {
		msg = "PLAYER 2 win"
	} else {
		msg = "Nobody win"
	}
	fmt.Println(msg)
}
