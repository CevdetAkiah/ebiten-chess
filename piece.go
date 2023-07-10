package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/notnil/chess"
)

// TODO: need to update the piece image if promoted

// a chess piece model
type Piece struct {
	location  chess.Square
	pieceType string
	side      string
	image     *ebiten.Image
}

// update the piece map with the new piece positions
func (g *Game) updatePieces(moveTo chess.Square) {
	// update the pieces map
	g.pieces[g.selectedPiece.location] = nil // remove the piece from the current location
	g.selectedPiece.location = moveTo
	g.pieces[moveTo] = g.selectedPiece // assign the piece to a new location

	g.selectedPiece = nil
}

// create a chess piece
func createPiece(piece chess.Piece, sq chess.Square) *Piece {
	if piece == chess.NoPiece {
		return nil
	}

	pType := piece.Type().String()
	colour := piece.Color().String()

	if colour == "w" {
		colour = "white"
	} else {
		colour = "black"
	}

	switch pType {
	case "r":
		pType = "rook"
	case "n":
		pType = "knight"
	case "q":
		pType = "queen"
	case "b":
		pType = "bishop"
	case "p":
		pType = "pawn"
	case "k":
		pType = "king"
	}

	chessPiece := &Piece{
		pieceType: pType,
		side:      colour,
		image:     nil,
		location:  sq,
	}
	setPieceImage(chessPiece, pType)
	return chessPiece
}

// set the piece image
func setPieceImage(p *Piece, pType string) {
	path := fmt.Sprintf("piecePNG/%s/%s.png", p.side, pType)

	pieceImage, _, _ := ebitenutil.NewImageFromFile(path)

	p.image = pieceImage
}
