package main

import (
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/notnil/chess"
)

type Game struct {
	playerTurn    string
	buttonPress   bool
	chessboard    *board
	gamestart     bool
	engine        *chess.Game
	pieces        map[chess.Square]*Piece
	selectedPiece *Piece
}

func (g *Game) Update() error {
	// detect mouse click
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		// ignore multiple inputs from the same click
		if g.buttonPress {
			g.buttonPress = false
			return nil
		}
		g.buttonPress = true

		// get mouse click position
		mouseX, mouseY := ebiten.CursorPosition()
		chessX, chessY := getChessGridCoordinates(mouseX, mouseY)

		// Make the move
		// if user has already selected a piece and the clicked square is valid, make the move
		if g.selectedPiece != nil {
			moveTo := squareOffset(chessX, chessY)
			if isValidSquare(g.chessboard.validSquares, moveTo) {
				// make the move in the engine and update the pieces map
				makeMove(g, moveTo)
				g.updatePieces(moveTo)

				// toggle the players turn on the board
				if g.playerTurn == "White" {
					g.playerTurn = "Black"
				} else {
					g.playerTurn = "White"
				}
				// update the pieces on the board
				g.updateBoardImage()
			}
		}

		// if mouse is clicked on a piece continue
		if piece := g.pieces[squareOffset(chessX, chessY)]; piece != nil {
			// turn off previously selected square or valid colour
			for _, cell := range g.chessboard.cells {
				cell.turnOffColours(g)
			}
			// mark the new cell as selected
			g.chessboard.cells[piece.location].selected = true
			g.updateCellColours(piece.location)
			g.selectedPiece = piece
		} else {
			g.buttonPress = false
		}

	}
	return nil
}

// re draw the board with the new piece positions
func (g *Game) updateBoardImage() {
	FEN := g.engine.Position().Board().SquareMap()
	for _, cell := range g.chessboard.cells {
		vector.DrawFilledRect(g.chessboard.grid, cell.xPos, cell.yPos, squareWidth, squareHeight, cell.squareColour, false)

		if piece := FEN[cell.position]; piece != chess.NoPiece {
			p := g.pieces[cell.position]
			g.positionPiece(p, cell)
		}
	}
}

// update the piece map with the new piece positions
func (g *Game) updatePieces(moveTo chess.Square) {
	// update the pieces map
	g.pieces[g.selectedPiece.location] = nil // remove the piece from the current location
	g.selectedPiece.location = moveTo
	g.pieces[moveTo] = g.selectedPiece // assign the piece to a new location
	g.selectedPiece = nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear the screen
	screen.Fill(color.RGBA{R: 50, G: 50, B: 50, A: 255}) // Set the screen background color to a gray

	// Draw the chessboard
	screen.DrawImage(g.chessboard.grid, nil)

	// Button to pick side
	if g.gamestart == true {
		// button to pick a side

		g.gamestart = false
	} else {
		// Draw label
		labelText := strings.ToUpper(g.playerTurn) + " PLAYING"
		font := LoadFont("./fonts/Roboto-Bold.ttf")
		label := newLabel(font, labelText, (screenWidth-measureTextWidth(font, labelText))/2, screenHeight+60)
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
