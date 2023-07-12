package main

import (
	"math/rand"
	"time"
)

// stand in for stockfish
func randomMove(g *Game) {
	time.Sleep(500 * time.Millisecond)
	rand.Seed(time.Now().Unix())
	// return if no side chosen
	if g.player == "" {
		return
	}
	valid := g.engine.Position().ValidMoves()
	m := valid[rand.Intn(len(valid))]
	s1 := m.S1()
	s2 := m.S2()
	g.selectedPiece = g.pieces[s1]
	g.updatePieces(s2)
	g.engine.Move(m)
	g.updateBoardImage()

	// toggle the players turn on the board
	if g.playerTurn == "White" {
		g.playerTurn = "Black"
	} else {
		g.playerTurn = "White"
	}
}
