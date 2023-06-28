package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	boardWidth   = 640
	boardHeight  = 480
	screenWidth  = 640
	screenHeight = 600
	rows         = 8
	columns      = 8
)

const (
	squareWidth  = boardWidth / columns
	squareHeight = boardHeight / rows
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("CHESS")

	game, err := NewGame()
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}
