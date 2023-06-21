package main

import (
	"github.com/notnil/chess"
)

func makeMove(g *Game, moveTo chess.Square) {
	moves := g.engine.Position().ValidMoves()
	for _, move := range moves {
		if move.S1() == g.selectedPiece.location && move.S2() == moveTo {
			g.engine.Move(move)
		}
	}
}
