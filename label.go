package main

import (
	"io/ioutil"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Label struct {
	font font.Face
	text string
	x    int
	y    int
}

func newLabel(font font.Face, labelString string, x int, y int) *Label {
	label := &Label{
		font: font,
		text: labelString,
		x:    x,
		y:    y,
	}
	return label
}

func LoadFont(path string) font.Face {
	fontBytes, err := ioutil.ReadFile(path)
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
