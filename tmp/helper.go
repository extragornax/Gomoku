package main

func newMove(x int, y int, poid int) Move {
	var tmp Move

	tmp.x = x
	tmp.y = y
	tmp.poid = poid
	tmp.pruned = false
	return tmp
}

func newLastChance(winner int, i int, j int, x int, y int) *LastChance {
	lastChance := new(LastChance)
	lastChance.winner = winner
	lastChance.i = i
	lastChance.j = j
	lastChance.x = x
	lastChance.y = y
	return lastChance
}

func copyGame(game *Game) *Game {
	var newGame *Game = new(Game)
	newGame.score = game.score
	newGame.lastChance = game.lastChance
	newGame.curPlayer = game.curPlayer
	newGame.board = game.board
	newGame.depth = game.depth + 1
	newGame.winner = game.winner
	newGame.friend = game.friend
	return newGame
}

func newGame(game *Game) {
	game.score[0] = 0
	game.score[1] = 0
	game.lastChance = nil
	game.curPlayer = PLAYER1
	game.friend = PLAYER1
	game.board = newBoard()
	game.depth = 0
	game.winner = 0
}

func initGame() *Game {
	var game *Game = new(Game)

	game.score[0] = 0
	game.score[1] = 0
	game.lastChance = nil
	game.curPlayer = PLAYER1
	game.friend = PLAYER1
	game.board = newBoard()
	game.depth = 0
	game.winner = 0
	return game
}

func newBoard() [SIZE][SIZE]int {
	var board [SIZE][SIZE]int
	var x = 0
	for x < SIZE {
		var y = 0
		for y < SIZE {
			board[x][y] = 0
			y++
		}
		x++
	}
	return board
}

func initTools() *Tools {
	tools := new(Tools)

	tools.iaStart = false
	tools.gameType = SOLO
	tools.wait = false
	tools.aiHelper = false
	return tools
}
