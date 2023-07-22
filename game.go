package main

import (
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/notnil/chess"
	"github.com/notnil/chess/uci"
)

type Game struct {
	player        string
	stockfish     *uci.Engine
	playerTurn    string
	buttonPress   bool
	chessboard    *board
	gamestart     bool
	engine        *chess.Game
	pieces        map[chess.Square]*Piece
	selectedPiece *Piece
	sb            []*startButton
	winner        string
}

func (g *Game) Update() error {
	outcome := g.engine.Outcome()
	// play until an outcome is decided
	if outcome == chess.NoOutcome {

		// detect mouse click
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			// ignore multiple inputs from the same click
			if g.buttonPress {
				g.buttonPress = false
				return nil
			}
			g.buttonPress = true
			mouseX, mouseY := ebiten.CursorPosition()
			if g.gamestart {
				awaitChoice(g, mouseX, mouseY)
			}
			// if player turn then play, else it's the ai's turn
			if g.playerTurn == g.player {
				// get mouse click position
				chessX, chessY := getChessGridCoordinates(mouseX, mouseY)
				moveTo := squareOffset(chessX, chessY)
				// make the move
				g.makeMove(moveTo)
				// update board colours
				g.updateBoardColours(chessX, chessY)
			} else {
				// random move // ai
				// randomMove(g)
				// stockfish move
				move := g.uciMove()
				g.selectedPiece = g.pieces[move.S1()]
				g.makeMove(move.S2())
			}

		}
	} else {
		switch outcome {
		case chess.WhiteWon:
			g.winner = "WHITE"
		case chess.BlackWon:
			g.winner = "BLACK"
		case chess.Draw:
			g.winner = "DRAW"
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear the screen
	screen.Fill(color.RGBA{R: 50, G: 50, B: 50, A: 255}) // Set the screen background color to a gray

	// Draw the chessboard
	screen.DrawImage(g.chessboard.grid, nil)

	font := LoadFont("static/fonts/Roboto-Bold.ttf")
	// Button to pick side
	if g.gamestart {
		chooseSide(g, screen, font)
	} else {
		// Draw label
		var labelText string
		if g.engine.Outcome() == chess.NoOutcome {
			labelText = strings.ToUpper(g.playerTurn) + " PLAYING"
		} else {
			declaration := " WINS"
			if g.winner == "DRAW" {
				declaration = ""
			}
			labelText = g.winner + declaration
		}
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
	sf := uciStockFish()
	game := &Game{
		player:     "",
		stockfish:  sf,
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
