package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/notnil/chess"
)

// returns a chess.Square
func squareOffset(column, row int) chess.Square {
	var sq chess.Square
	offset := row*8 + column

	sq = chess.Square(offset)
	return sq
}

// position the piece on the board according to the cell it occupies
func (g *Game) positionPiece(p *Piece, cell *cell) {
	// get the values needed to center the image on the cell
	imageWidth, imageHeight := p.image.Bounds().Dx(), p.image.Bounds().Dy()
	centerX := float64(cell.xPos) + float64(squareWidth)/2 - float64(imageWidth)/2
	centerY := float64(cell.yPos) + float64(squareHeight)/2 - float64(imageHeight)/2
	// Superimpose the p image on the square
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(centerX, centerY)
	g.chessboard.grid.DrawImage(p.image, opts)
}

func isValidSquare(squares []chess.Square, target chess.Square) bool {
	for _, sq := range squares {
		if sq == target {
			return true
		}
	}
	return false
}
