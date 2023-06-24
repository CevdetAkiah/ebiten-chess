package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
	rows         = 8
	columns      = 8
)

const (
	squareWidth  = screenWidth / columns
	squareHeight = screenHeight / rows
)

func main() {
	ebiten.SetWindowSize(640, 600)
	ebiten.SetWindowTitle("CHESS")

	game, err := NewGame()
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}
