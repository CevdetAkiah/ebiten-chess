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
	if c.selected {
		c.selected = false
		vector.DrawFilledRect(g.chessboard.grid, c.xPos, c.yPos, squareWidth, squareHeight, c.squareColour, false)
		piece := g.pieces[c.position]
		if piece != nil {
			g.positionPiece(piece, c)
		}
	}
	if c.valid {
		c.valid = false
		vector.DrawFilledRect(g.chessboard.grid, c.xPos, c.yPos, squareWidth, squareHeight, c.squareColour, false)
	}
}
