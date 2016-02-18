package main

import "github.com/JoelOtter/termloop"

type field struct {
	name  string
	value string
}

type InputDialog struct {
	*Dialog
	x        int
	y        int
	fg       termloop.Attr
	bg       termloop.Attr
	text     []string
	canvas   [][]termloop.Cell
	fields   []field
	original []string
}

func NewInputDialog(paddingX, paddingY int, text []string, fg, bg termloop.Attr) *InputDialog {
	// Find out text dimensions (adding padding)
	dialogWidth, dialogHeight := DialogDimensions(text, paddingX, paddingY)

	// Pad text
	text = PadText(text, paddingX, paddingY)

	// Center the dialog
	x, y := CenterDialog(dialogWidth, dialogHeight)

	canvas := make([][]termloop.Cell, dialogHeight)
	for i := range text {
		canvas[i] = make([]termloop.Cell, dialogWidth)
		for j := range canvas[i] {
			canvas[i][j] = termloop.Cell{Ch: ' ', Fg: fg, Bg: bg}
			if j < len(text[i]) {
				canvas[i][j].Ch = rune(text[i][j])
			}
		}
	}

	return &InputDialog{
		x:      x,
		y:      y,
		fg:     fg,
		bg:     bg,
		text:   text,
		canvas: canvas,
	}
}
