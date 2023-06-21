package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
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

		// if user has already selected a piece and the clicked square is valid, make the move
		if g.selectedPiece != nil {
			moveTo := squareOffset(chessX, chessY)
			if isValidSquare(g.chessboard.validSquares, moveTo) {
				makeMove(g, moveTo)
				// update the pieces map
				g.pieces[g.selectedPiece.location] = nil // remove the piece from the current location
				g.selectedPiece.location = moveTo
				g.pieces[moveTo] = g.selectedPiece // assign the piece to a new location
				g.selectedPiece = nil

				// update the board image
				FEN := g.engine.Position().Board().SquareMap()
				for _, cell := range g.chessboard.cells {
					vector.DrawFilledRect(g.chessboard.grid, cell.xPos, cell.yPos, squareWidth, squareHeight, cell.squareColour, false)

					if piece := FEN[cell.position]; piece != chess.NoPiece {
						p := g.pieces[cell.position]
						g.positionPiece(p, cell)
					}
				}
			}
		}

		// if mouse is clicked on a piece continue
		if piece := g.pieces[squareOffset(chessX, chessY)]; piece != nil {
			// turn off previously selected square or valid colour
			for _, cell := range g.chessboard.cells {
				if cell.selected {
					cell.selected = false
					vector.DrawFilledRect(g.chessboard.grid, cell.xPos, cell.yPos, squareWidth, squareHeight, cell.squareColour, false)
					piece := g.pieces[cell.position]
					if piece != nil {
						g.positionPiece(piece, cell)

					}
				}
				if cell.valid {
					cell.valid = false
					vector.DrawFilledRect(g.chessboard.grid, cell.xPos, cell.yPos, squareWidth, squareHeight, cell.squareColour, false)
				}
			}
			// mark the cell as selected
			g.chessboard.cells[int(piece.location)].selected = true
			g.updateCellColours(int(piece.location))
			g.selectedPiece = piece
		} else {
			g.buttonPress = false
		}

	}
	return nil
}

func getChessGridCoordinates(x, y int) (int, int) {
	column := x / squareWidth
	row := y / squareHeight
	return column, row
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear the screen
	screen.Fill(color.White)

	screen.DrawImage(g.chessboard.grid, nil)
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
