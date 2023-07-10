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

func (g *Game) makeMove(chessX, chessY int) {
	// if user has already selected a piece and the clicked square is valid, make the move
	if g.selectedPiece != nil {
		moveTo := squareOffset(chessX, chessY)
		if isValidSquare(g.chessboard.validSquares, moveTo) {
			// make the move in the engine and update the pieces map
			parseMoves(g, moveTo)
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

func isPromotionSquare(p *Piece, sq chess.Square) bool {
	if sq == chess.A8 || sq == chess.B8 || sq == chess.C8 || sq == chess.D8 || sq == chess.E8 || sq == chess.F8 || sq == chess.G8 || sq == chess.H8 ||
		sq == chess.A1 || sq == chess.B1 || sq == chess.C1 || sq == chess.D1 || sq == chess.E1 || sq == chess.F1 || sq == chess.G1 || sq == chess.H1 {
		if p.pieceType == "pawn" {
			return true
		}

	}

	return false
}
