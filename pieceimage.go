package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/notnil/chess"
)

func returnPieceImage(piece chess.Piece) *ebiten.Image {
	if piece == chess.NoPiece {
		return nil
	}
	pieceType := piece.Type().String()

	colour := piece.Color().String()

	if colour == "w" {
		colour = "white"
	} else {
		colour = "black"
	}

	switch pieceType {
	case "r":
		pieceType = "rook"
	case "n":
		pieceType = "knight"
	case "q":
		pieceType = "queen"
	case "b":
		pieceType = "bishop"
	case "p":
		pieceType = "pawn"
	case "k":
		pieceType = "king"
	}

	path := fmt.Sprintf("piecePNG/%s/%s.png", colour, pieceType)

	pieceImage, _, _ := ebitenutil.NewImageFromFile(path)

	return pieceImage
}
