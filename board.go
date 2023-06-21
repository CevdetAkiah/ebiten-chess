package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/notnil/chess"
)

type Game struct {
	buttonPress   bool
	chessboard    *board
	gamestart     bool
	engine        *chess.Game
	pieces        map[chess.Square]*Piece
	selectedPiece *Piece
}

type board struct {
	grid         *ebiten.Image
	validSquares []chess.Square
	cells        map[int]*cell
}

// each square on the grid held within the chessboard
type cell struct {
	squareColour color.Color
	selected     bool
	valid        bool
	position     chess.Square
	xPos         float32
	yPos         float32
}

// update the board after user inpt
func (g *Game) updateCellColours(sq int) {
	cell := g.chessboard.cells[sq]
	piece := g.pieces[chess.Square(sq)]
	// turn on the select colour for the selected cell and reposition the piece
	selectSquare(g, cell, piece)
	// turn on the valid square colours
	validSquare(g, cell, piece)
}

// show the valid squares on the board
func validSquare(g *Game, c *cell, p *Piece) {
	position := g.engine.Position()
	moves := position.ValidMoves()

	// reset valid squares
	g.chessboard.validSquares = nil

	// gather the valid moves for selected piece
	for _, move := range moves {
		if move.S1() == p.location {
			g.chessboard.validSquares = append(g.chessboard.validSquares, move.S2())
		}
	}

	// apply valid colour
	for _, sq := range g.chessboard.validSquares {
		cell := g.chessboard.cells[int(sq)]
		cell.valid = true
		colourSquare(cell, g, color.RGBA{R: 0, G: 255, B: 50, A: 80}) // valid square colour
	}
}

// turn on the selected square colour
func selectSquare(g *Game, c *cell, p *Piece) {
	colourSquare(c, g, color.RGBA{R: 181, G: 136, B: 99, A: 200}) // selected square colour
	g.positionPiece(p, c)
}

// apply the select colour to the selected cell
func colourSquare(c *cell, g *Game, colour color.Color) {
	originalColour := c.squareColour

	// set the cell to have the updated colour
	c.squareColour = colour
	vector.DrawFilledRect(g.chessboard.grid, c.xPos, c.yPos, squareWidth, squareHeight, c.squareColour, false)

	// reset the cell colour
	c.squareColour = originalColour
}

// start a new game
func NewGame() *Game {
	eng := chess.NewGame()

	game := &Game{
		gamestart: true,
		engine:    eng,
		pieces:    make(map[chess.Square]*Piece),
	}
	// set up the chessboard
	chessboard := createChessboard(game)
	game.chessboard = chessboard
	game.gamestart = true

	// place the pieces on the chessboard
	for _, piece := range game.pieces {
		cell := game.chessboard.cells[int(piece.location)]
		game.positionPiece(piece, cell)
	}

	return game

}

// draws the chessboard and loads the pieces into the game with their positions
func createChessboard(g *Game) *board {
	chessboard := &board{}
	chessboard.cells = make(map[int]*cell, rows*columns)

	// Create a new image for the grid
	grid := ebiten.NewImage(screenWidth, screenHeight)
	FEN := g.engine.Position().Board().SquareMap()
	chessboard.grid = grid
	// Draw the chess grid
	for row := 0; row < 8; row++ {
		for column := 0; column < 8; column++ {
			cell := &cell{}
			// Calculate the position of the square
			cell.xPos = float32(column * squareWidth)
			cell.yPos = float32(row * squareHeight)

			// Determine the colour of the square
			if (column+row)%2 == 0 {
				cell.squareColour = color.RGBA{R: 181, G: 136, B: 99, A: 255} // Light square colour
			} else {
				cell.squareColour = color.RGBA{R: 240, G: 217, B: 181, A: 255} // Dark square colour
			}

			// Draw the square
			vector.DrawFilledRect(chessboard.grid, cell.xPos, cell.yPos, squareWidth, squareHeight, cell.squareColour, false)

			// convert the position to chess square and configure cell
			fensq := squareOffset(column, row)
			cell.position = fensq
			// add cell to the map
			chessboard.cells[int(fensq)] = cell
			// Get the piece using the board position
			piece := FEN[fensq]

			// Draw the image on the square (if there is a piece on that square)
			if piece == chess.NoPiece {
				continue
			} else {
				// add the piece to the game struct for user interactivity
				piece := createPiece(piece, fensq)
				g.pieces[fensq] = piece
			}

		}
	}
	return chessboard
}
