package main

import "fmt"

func main() {
	game := &GameState{}
	go game.ProcessCSV()
	game.Init()
}
