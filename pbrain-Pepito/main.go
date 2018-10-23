package main

import (
	"gomoku/gomoku"
)

func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// 	s := scanner.Text()
	// 	fmt.Println(s)
	// }
	// if err := scanner.Err(); err != nil {
	// 	os.Exit(1)
	// }
	var game gomoku.Gomoku
	game.Init()
	game.Run()
	game.Play(2, 2)
}
