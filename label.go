package main

import (
	"image/color"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Label struct {
	font   font.Face
	text   string
	colour color.Color
	x      int
	y      int
}

func newLabel(font font.Face, labelString string, colour color.Color, x int, y int) *Label {
	label := &Label{
		font:   font,
		text:   labelString,
		colour: colour,
		x:      x,
		y:      y,
	}
	return label
}

func LoadFont(path string) font.Face {
	fontBytes, err := embeddedFiles.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	font, err := opentype.Parse(fontBytes)
	if err != nil {
		log.Fatal(err)
	}
	const fontSize = 12
	labelFont, err := opentype.NewFace(font, &opentype.FaceOptions{
		Size: fontSize,
		DPI:  100,
	})
	if err != nil {
		log.Fatal(err)
	}
	return labelFont
}

func measureTextWidth(face font.Face, text string) int {
	b, _ := font.BoundString(face, text)
	return (b.Max.X - b.Min.X).Ceil()
}
