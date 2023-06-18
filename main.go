package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/notnil/chess"
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

type Game struct {
	chessboard *ebiten.Image
	gamestart  bool
	engine     *chess.Game
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear the screen
	screen.Fill(color.White)

	screen.DrawImage(g.chessboard, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("CHESS")

	game := NewGame()

	go func() {
		if err := ebiten.RunGame(game); err != nil {
			log.Fatal(err)
		}
	}()

	select {}
}
