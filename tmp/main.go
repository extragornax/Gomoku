package main

// PLAYER1 const
const PLAYER1 = 1

// PLAYER2 const
const PLAYER2 = 2

// NUL const
const NUL = 3

// NONE const
const NONE = 0

// HINT const
const HINT = 3

var DEPTH_MAX int = 1

const EQUAL = 0
const INFINI = 1000000
const AMP = 1
const SIZE = 19
const MULTI = 2
const SOLO = 1
const MENU = 0

type Move struct {
	x      int
	y      int
	poid   int
	pruned bool
}

type AlphaBeta struct {
	alpha int
	beta  int
}

type LastChance struct {
	i      int
	j      int
	x      int
	y      int
	winner int
	player int
}

type Game struct {
	lastChance *LastChance
	score      [2]int
	curPlayer  int
	depth      int
	winner     int
	board      [SIZE][SIZE]int
	friend     int
}

type Tools struct {
	exit     bool
	iaStart  bool
	aiHelper bool
	gameType int
	wait     bool
}

func main() {
	tools := initTools()
	game := initGame()

	// displayMenu(tools, 0)
	handleEvent(tools, game)

	return
}
