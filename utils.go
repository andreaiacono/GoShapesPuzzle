package main

import (
	"github.com/gotk3/gotk3/cairo"
	"math/rand"
)


type Color struct {
	R float64
	G float64
	B float64
}

const EMPTY = 0
const FLOOD_FILL_VALUE = 255
const StartingSpeed = 500

func DrawRectangle(x float64, y float64, width float64, height float64, cr *cairo.Context, num string) {
	cr.MoveTo(x, y)
	cr.LineTo(x+width, y)
	cr.LineTo(x+width, y+height)
	cr.LineTo(x, y+height)
	cr.LineTo(x, y)
	cr.Stroke()
	if num != "" {
		cr.MoveTo(x+10, y+25)
		setColor(Color{0,0,0}, cr)
		cr.SetFontSize(float64(15))
		cr.SelectFontFace("Arial", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_NORMAL)
		cr.ShowText(num)
	}
}


func GenerateColors(n int) []Color {
	rand.Seed(201712)
	var colors = make([]Color, n)
	var i int
	for i = 0; i < n; i++ {
		colors[i] = Color{rand.Float64(), rand.Float64(), rand.Float64()}
	}
	return colors[:]
}
