package main

import (
	"github.com/notnil/chess"
)

func parseMoves(g *Game, moveTo chess.Square) {
	moves := g.engine.Position().ValidMoves()
	for _, move := range moves {
		if move.S1() == g.selectedPiece.location && move.S2() == moveTo {
			g.engine.Move(move)
		}
	}
}

func (g *Game) makeMove(moveTo chess.Square) {
	// if user has already selected a piece and the clicked square is valid, make the move
	if g.selectedPiece != nil {
		if isValidSquare(g.chessboard.validSquares, moveTo) {
			// make the move in the engine and update the pieces map
			parseMoves(g, moveTo)
			// if pawn to promote then update the pawn image
			if isPromotionSquare(g.selectedPiece, moveTo) {
				setPieceImage(g.selectedPiece, "queen")
			}
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
}
