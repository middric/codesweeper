package main

import (
	"strconv"
	"strings"

	"github.com/JoelOtter/termloop"
)

type Dialog struct {
	x        int
	y        int
	fg       termloop.Attr
	bg       termloop.Attr
	text     []string
	original []string
	canvas   [][]termloop.Cell
}

// Determine length of longest string in a slice of strings
func maximumLength(text []string) int {
	length := 0

	for i := range text {
		if len(text[i]) > length {
			length = len(text[i])
		}
	}

	return length
}

// DialogDimensions takes a slice of strings and an amount of padding and
// determines the correct width and height for a dialog box
func DialogDimensions(text []string, paddingX int, paddingY int) (int, int) {
	dialogWidth := maximumLength(text) + (paddingX * 2)
	dialogHeight := len(text) + (paddingY * 2)

	return dialogWidth, dialogHeight
}

// PadText iterates through a slice of strings and adds the appropriate top and
// side padding
func PadText(text []string, paddingX, paddingY int) []string {
	text = append(text, make([]string, paddingY)...)
	text = append(make([]string, paddingY), text...)
	for i := range text {
		// Pad line
		text[i] = strings.Repeat(" ", paddingX) + text[i] + strings.Repeat(" ", paddingX)
	}
	return text
}

// CenterDialog takes a dialog width and height and return the x and y
// coordinates to use when centering the dialog
func CenterDialog(dialogWidth, dialogHeight int) (int, int) {
	x := (width - dialogWidth) / 2
	y := (height - dialogHeight) / 2
	return x, y
}

func isInputField(text string) bool {
	trimmed := strings.TrimSpace(text)
	prefix := strings.HasPrefix(trimmed, "{{")
	suffix := strings.HasSuffix(trimmed, "}}")

	return prefix && suffix
}

func inputKeyValue(input string) (string, string) {
	input = strings.TrimSpace(input)
	input = strings.TrimPrefix(input, "{{")
	input = strings.TrimSuffix(input, "}}")

	components := strings.Split(input, ",")

	return strings.TrimSpace(components[0]), strings.TrimSpace(components[1])
}

// NewDialog creates a new Dialog object
func NewDialog(paddingX, paddingY int, text []string, fg, bg termloop.Attr) *Dialog {
	original := text
	// Find out text dimensions (adding padding)
	dialogWidth, dialogHeight := DialogDimensions(text, paddingX, paddingY)

	// Pad text
	text = PadText(text, paddingX, paddingY)

	// Center the dialog
	x, y := CenterDialog(dialogWidth, dialogHeight)

	canvas := make([][]termloop.Cell, dialogHeight)
	for i := range text {
		canvas[i] = make([]termloop.Cell, dialogWidth)
		if isInputField(text[i]) {
			canvas[i][0] = termloop.Cell{Ch: ' ', Fg: fg, Bg: bg}
			for j := 1; j < dialogWidth-1; j++ {
				canvas[i][j] = termloop.Cell{Ch: '_', Fg: fg, Bg: bg}
			}
			_, value := inputKeyValue(text[i])
			for j := 0; j < len(value); j++ {
				canvas[i][len(canvas[i])-2-j].Ch = rune(value[len(value)-1-j])
			}
			canvas[i][dialogWidth-1] = termloop.Cell{Ch: ' ', Fg: fg, Bg: bg}
		} else {
			for j := range canvas[i] {
				canvas[i][j] = termloop.Cell{Ch: ' ', Fg: fg, Bg: bg}
				if j < len(text[i]) {
					canvas[i][j].Ch = rune(text[i][j])
				}
			}
		}
	}

	return &Dialog{
		x:        x,
		y:        y,
		fg:       fg,
		bg:       bg,
		text:     text,
		canvas:   canvas,
		original: original,
	}
}

func (d *Dialog) Close() {
	level.RemoveEntity(d)
	player.state = Alive
}

func (d *Dialog) Tick(ev termloop.Event) {
	// Key presses between 0-9
	if ev.Ch >= 48 && ev.Ch <= 57 {
		mineCount, _ = strconv.Atoi(string(ev.Ch))
	} else if ev.Type == termloop.EventKey {
		switch ev.Key {
		case termloop.KeyEnter:
		case termloop.KeyEsc:
			d.Close()
			break
		}
	}
}

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
