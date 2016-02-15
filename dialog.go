package main

import (
	"strings"

	"github.com/JoelOtter/termloop"
)

type Dialog struct {
	x      int
	y      int
	fg     termloop.Attr
	bg     termloop.Attr
	text   []string
	canvas [][]termloop.Cell
}

func longest(text []string) int {
	length := 0

	for i := range text {
		if len(text[i]) > length {
			length = len(text[i])
		}
	}

	return length
}

func NewDialog(paddingX, paddingY int, text []string, fg, bg termloop.Attr) *Dialog {

	// Find out text dimensions (adding padding)
	dialogWidth := longest(text) + (paddingX * 2)
	dialogHeight := len(text) + (paddingY * 2)

	// Pad dialog with a blank line before and after
	text = append(text, make([]string, paddingY)...)
	text = append(make([]string, paddingY), text...)

	canvas := make([][]termloop.Cell, dialogHeight)
	for i, line := range text {
		// Pad line
		str := []rune(strings.Repeat(" ", paddingX) + line + strings.Repeat(" ", paddingX))

		canvas[i] = make([]termloop.Cell, dialogWidth)

		for j := range canvas[i] {
			ch := ' '
			if j < len(str) {
				ch = str[j]
			}
			canvas[i][j] = termloop.Cell{Ch: ch, Fg: fg, Bg: bg}
		}
	}

	// Center the dialog
	x := (width - dialogWidth) / 2
	y := (height - dialogHeight) / 2

	return &Dialog{
		x:      x,
		y:      y,
		fg:     fg,
		bg:     bg,
		text:   text,
		canvas: canvas,
	}
}

func (d *Dialog) Tick(ev termloop.Event) {}

// Draw draws the Dialog to the Screen s.
func (d *Dialog) Draw(s *termloop.Screen) {
	width, height := d.Size()

	// Render shadow
	for x := 1; x < width; x++ {
		s.RenderCell(d.x+x, d.y+height, &termloop.Cell{Ch: '▒', Fg: d.fg, Bg: termloop.ColorDefault})
	}
	for y := 1; y < height; y++ {
		s.RenderCell(d.x+width, d.y+y, &termloop.Cell{Ch: '▒', Fg: d.fg, Bg: termloop.ColorDefault})
	}
	s.RenderCell(d.x+width, d.y+height, &termloop.Cell{Ch: '▒', Fg: d.fg, Bg: termloop.ColorDefault})

	// Render content
	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			s.RenderCell(d.x+y, d.y+x, &d.canvas[x][y])
		}
	}
}

// Position returns the (x, y) coordinates of the Text.
func (d *Dialog) Position() (int, int) {
	return d.x, d.y
}

// Size returns the width and height of the Text.
func (d *Dialog) Size() (int, int) {
	return len(d.canvas[0]), len(d.canvas)
}
