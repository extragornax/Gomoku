package gomoku

import (
	"fmt"
	"strconv"
	"strings"
)

type messageProcessor func(*Gomoku, []string)

var messageProcessorsMap = map[string]messageProcessor{
	"START":   messageProcessorStart,
	"TURN":    messageProcessorTurn,
	"BEGIN":   messageProcessorBegin,
	"BOARD":   messageProcessorBoard,
	"ABOUT":   messageProcessorAbout,
	"END":     messageProcessorEnd,
	"RESTART": messageProcessorRestart,
}

func messageProcessorStart(gmk *Gomoku, msg []string) {
	nb, _ := strconv.Atoi(msg[0])
	err := gmk.Board.Init(uint(nb))
	if err != nil {
		fmt.Println("ERROR " + err.Error())
		gmk.Live = false
	} else {
		fmt.Println("OK")
	}
}

func messageProcessorTurn(gmk *Gomoku, msg []string) {
	splited := strings.Split(msg[0], ",")
	x, _ := strconv.Atoi(splited[0])
	y, _ := strconv.Atoi(splited[1])
	gmk.Board.Cells[y][x] = BoardCellFoe
	gmk.Turn = true
}

func messageProcessorBegin(gmk *Gomoku, msg []string) {
	gmk.Begin = true
	gmk.Turn = true
}

func messageProcessorBoard(gmk *Gomoku, msg []string) {
	err := gmk.Board.Init(gmk.Board.Size)
	if err != nil {
		fmt.Println("ERROR " + err.Error())
		return
	}
	gmk.scanner.Scan()
	line := gmk.scanner.Text()
	for line != "DONE" && line != "" {
		messageProcessorBoardLine(gmk, line)
		gmk.scanner.Scan()
		line = gmk.scanner.Text()
	}
	gmk.Turn = true
}

func messageProcessorBoardLine(gmk *Gomoku, line string) {
	words := strings.Split(line, ",")
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

func messageProcessorEnd(gmk *Gomoku, msg []string) {
	// gmk.Live = false
}

func messageProcessorRestart(gmk *Gomoku, msg []string) {
	gmk.Live = true
	gmk.Turn = false
	gmk.Begin = false
}
