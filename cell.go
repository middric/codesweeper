package main

import (
	"strconv"

	"github.com/JoelOtter/termloop"
)

type Cell struct {
	proximity int
	isMine    bool
	entity    *termloop.Text
	render    bool
	isFlagged bool
}

func (cell *Cell) Draw(screen *termloop.Screen) {
	if cell.render {
		cell.entity.SetColor(termloop.ColorWhite, termloop.ColorBlue)
		if cell.isMine {
			cell.entity.SetText(sMine)
			cell.entity.SetColor(termloop.ColorRed, termloop.ColorBlack)
		} else if cell.proximity > 0 {
			cell.entity.SetText(strconv.Itoa(cell.proximity))
			cell.entity.SetColor(termloop.ColorYellow, termloop.ColorBlue)
		} else if cell.entity.Text() == sWave {
			cell.entity.SetText(sBubble)
		}
	} else if cell.isFlagged {
		cell.entity.SetText(sSpace)
		cell.entity.SetColor(termloop.ColorMagenta, termloop.ColorMagenta)
	} else {
		char := sSpace
		x, y := cell.entity.Position()
		if (x%2 == 0 && y%2 == 0) || (x%2 != 0 && y%2 != 0) {
			char = sWave
		}
		cell.entity.SetText(char)
		cell.entity.SetColor(termloop.ColorWhite, termloop.ColorCyan)
	}
	cell.entity.Draw(screen)
}

func (cell *Cell) Tick(event termloop.Event) {}
