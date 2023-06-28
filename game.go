package main

import (
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/notnil/chess"
)

type Game struct {
	player        string
	playerTurn    string
	buttonPress   bool
	chessboard    *board
	gamestart     bool
	engine        *chess.Game
	pieces        map[chess.Square]*Piece
	selectedPiece *Piece
	sb            []*startButton
}

func (g *Game) Update() error {
	// detect mouse click
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mouseX, mouseY := ebiten.CursorPosition()

		if g.gamestart {
			awaitChoice(g, mouseX, mouseY)
		}
		// ignore multiple inputs from the same click
		if g.buttonPress {
			g.buttonPress = false
			return nil
		}
		g.buttonPress = true

		// get mouse click position
		chessX, chessY := getChessGridCoordinates(mouseX, mouseY)

		// make the move
		g.makeMove(chessX, chessY)

		// update board colours
		g.updateBoardColours(chessX, chessY)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear the screen
	screen.Fill(color.RGBA{R: 50, G: 50, B: 50, A: 255}) // Set the screen background color to a gray

	// Draw the chessboard
	screen.DrawImage(g.chessboard.grid, nil)

	font := LoadFont("./fonts/Roboto-Bold.ttf")
	// Button to pick side
	if g.gamestart {
		chooseSide(g, screen, font)
	} else {
		// Draw label
		labelText := strings.ToUpper(g.playerTurn) + " PLAYING"
		label := newLabel(font, labelText, color.White, (boardWidth-measureTextWidth(font, labelText))/2, boardHeight+60)
		text.Draw(screen, labelText, label.font, label.x, label.y, color.White)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

// start a new game
func NewGame() (*Game, error) {
	eng := chess.NewGame()
	game := &Game{
		player:     "",
		playerTurn: "White",
		gamestart:  true,
		engine:     eng,
		pieces:     make(map[chess.Square]*Piece),
	}
	// set up the chessboard
	chessboard := createChessboard(game)
	game.chessboard = chessboard
	game.gamestart = true

	// place the pieces on the chessboard
	for _, piece := range game.pieces {
		cell := game.chessboard.cells[piece.location]
		game.positionPiece(piece, cell)
	}

	return game, nil

}
