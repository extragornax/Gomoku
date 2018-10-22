package gomoku

import (
	"fmt"
	"strconv"
	"strings"
)

// *** Message processors ***
type messageProcessor func(*Gomoku, []string)

var messagesMap = map[string]messageProcessor{
	"BOARD": processorBoard,
}

func processorBoard(gmk *Gomoku, msg []string) {
	err := gmk.Board.init(gmk.Board.Size)
	if err != nil {
		return
	}
	var line string
	fmt.Scanln(&line)
	for line != "DONE" {
		processorBoardLine(gmk, line)
		fmt.Scanln(&line)
	}
}

func processorBoardLine(gmk *Gomoku, line string) {
	words := strings.Fields(line)
	if len(words) == 3 {
		x, _ := strconv.Atoi(words[0])
		y, _ := strconv.Atoi(words[1])
		field, _ := strconv.Atoi(words[2])
		gmk.Board.Cells[y][x] = uint8(field)
	}
}

// *** Gomoku Object ***

// Gomoku main object
type Gomoku struct {
	Board Board
	Live  bool
}

// Init initialises the gomoku object
func (gmk *Gomoku) Init() {
	gmk.Live = true
}

// Run is the main exectuion loop
func (gmk *Gomoku) Run() {
	var message string
	for gmk.Live {
		fmt.Println("DEBUG expecting input from Piskvork")
		fmt.Scanln(&message)
		words := strings.Fields(message)
		if len(words) != 0 {
			call := messagesMap[words[0]]
			if call != nil {
				call(gmk, words[1:])
			}
		}
	}
}
