package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func play(game *Game, tools *Tools, i int, j int) {
	simulateMove(game, i, j)
	displayBoard(tools, game)
}

// IA Play
func iaTurn(tools *Tools, game *Game) time.Duration {
	alpha, beta := -INFINI, INFINI
	tools.wait = true
	timeBfore := time.Now()
	bestMove := getNextMove(game, alpha, beta)
	timeAfter := time.Now()
	time := timeAfter.Sub(timeBfore)
	play(game, tools, bestMove.x, bestMove.y)
	tools.wait = false
	return time
}

// AI search and display the Hint (green pion)
func displayHint(tools *Tools, game *Game) {
	tools.wait = true
	alpha, beta := -INFINI, INFINI
	bestMove := getNextMove(game, alpha, beta)

	game.board[bestMove.x][bestMove.y] = HINT
	displayBoard(tools, game)
	game.board[bestMove.x][bestMove.y] = NONE
	tools.wait = false
}

// Called when click of the screen
func onClic(i int, j int, tools *Tools, game *Game) {

	// If move is playable
	if isPlayable(game, i, j) {
		// Play
		play(game, tools, i, j)
		// If winner, end the game
		if game.winner != NONE {
			endGame(tools, game)
			// If game was against IA, make it play
		} else if tools.gameType == SOLO {
			time := iaTurn(tools, game)
			if game.winner != NONE {
				endGame(tools, game)
			} else if tools.aiHelper {
				displayHint(tools, game)
			}
			displayTime(tools, time)
		} else if tools.aiHelper {
			displayHint(tools, game)
		}
	} else {
		fmt.Println("Error: ", i, " ", j, " is not playable")
	}
}

func handleEvent(tools *Tools, game *Game) {

	reader := bufio.NewReader(os.Stdin)
	for !tools.exit {
		fmt.Println("Enter text: ")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
		words := strings.Fields(text)
		one, _ := strconv.Atoi(words[0])
		two, _ := strconv.Atoi(words[1])
		onClic(one, two, tools, game)
	}

}
