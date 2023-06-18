package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/notnil/chess"
)

func NewGame() *Game {
	eng := chess.NewGame()
	chessboardImage := createChessboard(eng)

	game := &Game{
		chessboard: chessboardImage,
		gamestart:  true,
		engine:     eng,
	}

	return game

}

func createChessboard(engine *chess.Game) *ebiten.Image {
	// Create a new image for the chessboard
	chessboard := ebiten.NewImage(screenWidth, screenHeight)
	FEN := engine.Position().Board().SquareMap()

	// Draw the chess grid
	for row := 0; row < 8; row++ {
		for column := 0; column < 8; column++ {
			// Calculate the position of the square
			squareX := column * squareWidth
			squareY := row * squareHeight

			// Determin the colour of the square
			var squareColour color.Color
			if (column+row)%2 == 0 {
				squareColour = color.RGBA{R: 181, G: 136, B: 99, A: 255} // Light square colour
			} else {
				squareColour = color.RGBA{R: 240, G: 217, B: 181, A: 255} // Dark square colour
			}

			// Draw the square
			vector.DrawFilledRect(chessboard, float32(squareX), float32(squareY), squareWidth, squareHeight, squareColour, false)

			// Get the piece using the board position
			fensq := squareOffset(column, row)
			piece := FEN[fensq]

			// Draw the image on the square (if there is a piece on that square)
			if piece == chess.NoPiece {
				continue
			} else {
				pieceImage := returnPieceImage(piece)

				// Position the image in the center of the square
				imageWidth, imageHeight := pieceImage.Bounds().Dx(), pieceImage.Bounds().Dy()
				centerX := float64(squareX) + float64(squareWidth)/2 - float64(imageWidth)/2
				centerY := float64(squareY) + float64(squareHeight)/2 - float64(imageHeight)/2
				// Superimpose the piece image on the square
				opts := &ebiten.DrawImageOptions{}
				opts.GeoM.Translate(centerX, centerY)
				chessboard.DrawImage(pieceImage, opts)
			}

		}
	}
	return chessboard
}

func squareOffset(column, row int) chess.Square {
	var sq chess.Square
	offset := row*8 + column

	sq = chess.Square(offset)
	return sq
}
