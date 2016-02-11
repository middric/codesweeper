package main

import (
	"strconv"

	"github.com/JoelOtter/termloop"
)

// Cell defines the properties for a cell in the minefield
type Cell struct {
	entity     *termloop.Text
	isFlagged  bool
	isMine     bool
	isRevealed bool
	isWave     bool
	proximity  int
}

// NewCell creates a new cell. Accepts x and y coordinate parameters to
// determine "isWave" state. Returns a new instance of Cell.
func NewCell(x int, y int) Cell {
	cell := Cell{
		entity:     termloop.NewText(x, y, " ", termloop.ColorWhite, termloop.ColorCyan),
		isFlagged:  false,
		isMine:     false,
		isWave:     (x%2 == 0 && y%2 == 0) || (x%2 != 0 && y%2 != 0),
		proximity:  0,
		isRevealed: false,
	}

	return cell
}

// Draw a cell
func (cell *Cell) Draw(screen *termloop.Screen) {
	if cell.isRevealed && cell.isMine {
		cell.drawMine()
	} else if cell.isRevealed {
		cell.drawRevealed()
	} else if cell.isFlagged {
		cell.drawFlag()
	} else {
		cell.drawHidden()
	}

	cell.entity.Draw(screen)
}

// Tick handles Cell inputs
func (cell *Cell) Tick(event termloop.Event) {}

func (cell *Cell) drawRevealed() {
	if cell.proximity > 0 {
		cell.drawProximity()
	} else {
		char := " "
		if cell.isWave {
			char = "·"
		}
		cell.entity.SetText(char)
		cell.entity.SetColor(termloop.ColorWhite, termloop.ColorBlue)
	}
}

func (cell *Cell) drawMine() {
	cell.entity.SetText("✱")
	cell.entity.SetColor(termloop.ColorRed, termloop.ColorBlack)
}

func (cell *Cell) drawProximity() {
	cell.entity.SetText(strconv.Itoa(cell.proximity))
	cell.entity.SetColor(termloop.ColorYellow, termloop.ColorBlue)
}

func (cell *Cell) drawFlag() {
	cell.entity.SetText("✖")
	cell.entity.SetColor(termloop.ColorRed, termloop.ColorWhite)
}

func (cell *Cell) drawHidden() {
	char := " "
	if cell.isWave {
		char = "~"
	}
	cell.entity.SetText(char)
	cell.entity.SetColor(termloop.ColorWhite, termloop.ColorCyan)
}
