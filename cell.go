package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/notnil/chess"
)

// each square on the grid held within the chessboard
type cell struct {
	squareColour color.Color
	selected     bool
	valid        bool
	position     chess.Square
	xPos         float32
	yPos         float32
}

func (c *cell) turnOffColours(g *Game) {
	piece := g.pieces[c.position]
	if c.selected {
		c.selected = false
		vector.DrawFilledRect(g.chessboard.grid, c.xPos, c.yPos, squareWidth, squareHeight, c.squareColour, false)
		if piece != nil {
			g.positionPiece(piece, c)
		}
	}
	if c.valid {
		c.valid = false
		vector.DrawFilledRect(g.chessboard.grid, c.xPos, c.yPos, squareWidth, squareHeight, c.squareColour, false)
		if piece != nil {
			g.positionPiece(piece, c)
		}
	}
}

// update the select or valid colours in a cell
func (g *Game) updateBoardColours(chessX, chessY int) {
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
