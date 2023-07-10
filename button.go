package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
)

type startButton struct {
	x      int
	y      int
	width  int
	height int
	colour color.Color
	label  *Label
	chosen bool
}

func awaitChoice(g *Game, mouseX, mouseY int) {
	g.sb[0].detectChoice(mouseX, mouseY) // white button
	g.sb[1].detectChoice(mouseX, mouseY) // black button
	if g.sb[0].chosen {
		g.player = "White"
		g.gamestart = false
	} else if g.sb[1].chosen {
		g.player = "Black"
		g.gamestart = false
	}
}

// detect button click within the button's bounds
func (b *startButton) detectChoice(mouseX, mouseY int) {
	if mouseX < b.width+b.x && mouseX > b.x && mouseY > boardHeight {
		b.chosen = true
	}
}

func chooseSide(g *Game, screen *ebiten.Image, font font.Face) {
	whiteStartButton := newStartButton(0, boardHeight, boardWidth/2, screenHeight-boardHeight, color.White, "WHITE", font, color.Black)
	blackStartButton := newStartButton(boardWidth/2, boardHeight, boardWidth/2, screenHeight-boardHeight, color.Black, "BLACK", font, color.White)
	g.sb = append(g.sb, whiteStartButton)
	g.sb = append(g.sb, blackStartButton)
	whiteStartButton.drawButton(screen)
	blackStartButton.drawButton(screen)
}

func newStartButton(bx, by, bwidth, bheight int, bcolour color.Color, labelText string, font font.Face, labelColour color.Color) *startButton {
	button := &startButton{}
	button.x = bx
	button.y = by
	button.width = bwidth
	button.height = bheight
	button.colour = bcolour

	// Calculate label	 position
	labelWidth := measureTextWidth(font, labelText)
	labelX := bx + (bwidth-labelWidth)/2
	labelY := by + ((screenHeight-boardHeight)+font.Metrics().Height.Ceil())/2

	button.label = newLabel(font, labelText, labelColour, labelX, labelY)
	return button
}

func (sb *startButton) drawButton(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(sb.x), float32(sb.y), float32(sb.width), float32(sb.height), sb.colour, false)
	text.Draw(screen, sb.label.text, sb.label.font, sb.label.x, sb.label.y, sb.label.colour)
}
