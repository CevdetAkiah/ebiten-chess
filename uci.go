package main

import (
	"os/exec"
	"path/filepath"
	"time"

	"github.com/notnil/chess"
	"github.com/notnil/chess/uci"
)

func uciStockFish() *uci.Engine {
	// set up engine to use stockfish exe

	stockfishPath, err := filepath.Abs("/usr/games/stockfish")
	if err != nil {
		panic(err)
	}

	cmd := exec.Command(stockfishPath)
	stockfish, err := uci.New(cmd.Path)
	if err != nil {
		panic(err)
	}
	// initialize uci with new game
	if err := stockfish.Run(uci.CmdUCI, uci.CmdIsReady, uci.CmdUCINewGame); err != nil {
		panic(err)
	}

	return stockfish
}

func (g *Game) uciMove() *chess.Move {
	time.Sleep(time.Millisecond * 250)
	cmdPos := uci.CmdPosition{Position: g.engine.Position()}
	cmdGo := uci.CmdGo{MoveTime: time.Second / 100}
	if err := g.stockfish.Run(cmdPos, cmdGo); err != nil {
		panic(err)
	}
	move := g.stockfish.SearchResults().BestMove
	piece := g.pieces[move.S1()]
	moves := g.engine.Position().ValidMoves()
	// gather the valid squares for the selected piece
	g.chessboard.validSquares = appendValidSquares(g.chessboard.validSquares, moves, piece)
	return move
}
