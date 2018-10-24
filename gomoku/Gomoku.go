package gomoku

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// Gomoku main object
type Gomoku struct {
	Board   Board
	Live    bool
	Turn    bool
	Begin   bool
	scanner *bufio.Scanner
}

// Init initialises the gomoku object
func (gmk *Gomoku) Init() {
	gmk.Live = true
	gmk.Turn = false
	gmk.Begin = false
	gmk.scanner = bufio.NewScanner(os.Stdin)
}

// Run is the main exectuion loop
func (gmk *Gomoku) Run() {
	gmk.Turn = false
	for gmk.Live && !gmk.Turn {
		fmt.Println("DEBUG expecting input from Piskvork")
		if gmk.scanner.Scan() == false {
			break
		}
		message := gmk.scanner.Text()
		fmt.Println("DEBUG Recieved: " + message)
		words := strings.Fields(message)
		if len(words) != 0 {
			call := messageProcessorsMap[words[0]]
			if call != nil {
				call(gmk, words[1:])
			} else {
				fmt.Println("DEBUG Don't know what to do with: " + words[0])
			}
		}
	}
}

// Play allows the brain to play on a tile
func (gmk *Gomoku) Play(x uint, y uint) error {
	if x >= gmk.Board.Size || y >= gmk.Board.Size {
		return errors.New("out of range")
	}
	if gmk.Board.Cells[y][x] == BoardCellEmpty {
		gmk.Board.Cells[y][x] = BoardCellOwn
		fmt.Printf("%d,%d\n", x, y)
	} else {
		return errors.New("Cell is not available")
	}
	return nil
}
