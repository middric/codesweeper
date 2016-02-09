package main

import (
	"strconv"

	"github.com/JoelOtter/termloop"
)

type Cell struct {
	entity    *termloop.Text
	isFlagged bool
	isMine    bool
	isWave    bool
	proximity int
	render    bool
}

func NewCell(x int, y int) Cell {
	cell := Cell{
		entity:    termloop.NewText(x, y, " ", termloop.ColorWhite, termloop.ColorCyan),
		isFlagged: false,
		isMine:    false,
		isWave:    (x%2 == 0 && y%2 == 0) || (x%2 != 0 && y%2 != 0),
		proximity: 0,
		render:    false,
	}

	return cell
}

func (cell *Cell) Draw(screen *termloop.Screen) {
	if cell.render {
		if cell.isMine {
			cell.drawMine()
		} else if cell.proximity > 0 {
			cell.drawProximity()
		} else {
			cell.drawRevealed()
		}
	} else if cell.isFlagged {
		cell.drawFlag()
	} else {
		cell.drawHidden()
	}

	cell.entity.Draw(screen)
}

func (cell *Cell) Tick(event termloop.Event) {}

func (cell *Cell) drawRevealed() {
	char := " "
	if cell.isWave {
		char = "·"
	}
	cell.entity.SetText(char)
	cell.entity.SetColor(termloop.ColorWhite, termloop.ColorBlue)
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
	cell.entity.SetText(sSpace)
	cell.entity.SetColor(termloop.ColorMagenta, termloop.ColorMagenta)
}

func (cell *Cell) drawHidden() {
	char := " "
	if cell.isWave {
		char = "~"
	}
	cell.entity.SetText(char)
	cell.entity.SetColor(termloop.ColorWhite, termloop.ColorCyan)
}
