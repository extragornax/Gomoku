package gomoku

import (
	"fmt"
	"strconv"
	"strings"
)

type messageProcessor func(*Gomoku, []string)

var messageProcessorsMap = map[string]messageProcessor{
	"START": messageProcessorStart,
	"TURN":  messageProcessorTurn,
	"BEGIN": messageProcessorBegin,
	"BOARD": messageProcessorBoard,
	"ABOUT": messageProcessorAbout,
}

func messageProcessorStart(gmk *Gomoku, msg []string) {
	nb, _ := strconv.Atoi(msg[0])
	err := gmk.Board.init(uint(nb))
	if err != nil {
		fmt.Println("ERROR " + err.Error())
		gmk.Live = false
	}
}

func messageProcessorTurn(gmk *Gomoku, msg []string) {
	x, _ := strconv.Atoi(msg[0])
	y, _ := strconv.Atoi(msg[1])
	gmk.Board.Cells[y][x] = BoardCellFoe
	gmk.Turn = true
}

func messageProcessorBegin(gmk *Gomoku, msg []string) {
	gmk.Turn = true
}

func messageProcessorBoard(gmk *Gomoku, msg []string) {
	err := gmk.Board.init(gmk.Board.Size)
	if err != nil {
		return
	}
	var line string
	fmt.Scanln(&line)
	for line != "DONE" {
		messageProcessorBoardLine(gmk, line)
		fmt.Scanln(&line)
	}
}

func messageProcessorBoardLine(gmk *Gomoku, line string) {
	words := strings.Fields(line)
	if len(words) == 3 {
		x, _ := strconv.Atoi(words[0])
		y, _ := strconv.Atoi(words[1])
		field, _ := strconv.Atoi(words[2])
		gmk.Board.Cells[y][x] = uint8(field)
	}
}

func messageProcessorAbout(gmk *Gomoku, msg []string) {
	fmt.Println("name=\"pepito\", version=\"1.0\", author=\"pepito\", country=\"China\"")
}
