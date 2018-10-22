package gomoku

import (
	"fmt"
	"strings"
)

// Gomoku main object
type Gomoku struct {
	Board Board
	Live  bool
	Turn  bool
}

// Init initialises the gomoku object
func (gmk *Gomoku) Init() {
	gmk.Live = true
	gmk.Turn = false
}

// Run is the main exectuion loop
func (gmk *Gomoku) Run() {
	var message string
	for gmk.Live && !gmk.Turn {
		fmt.Println("DEBUG expecting input from Piskvork")
		fmt.Scanln(&message)
		words := strings.Fields(message)
		if len(words) != 0 {
			call := messageProcessorsMap[words[0]]
			if call != nil {
				call(gmk, words[1:])
			}
		}
	}
}
