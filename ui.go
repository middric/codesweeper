package main

import (
	"strconv"

	"github.com/JoelOtter/termloop"
)

func SetupUI() {
	game.Screen().AddEntity(termloop.NewText(0, 0, " MineSweeper ", termloop.ColorBlue, termloop.ColorBlack))
}

func UpdateUI() {
	_, screenH := game.Screen().Size()
	game.Screen().AddEntity(termloop.NewText(0, screenH-1, " Mines: "+strconv.Itoa(mineCount)+", Flags: "+strconv.Itoa(flags), termloop.ColorBlue, termloop.ColorBlack))
}

func ShowGameOver() {
	msg := []string{
		"                     ",
		"      Game Over!     ",
		"                     ",
	}
	length := len(msg[0])
	left := (width / 2) - (length / 2)
	top := (height / 2) - 2

	for i, line := range msg {
		level.AddEntity(termloop.NewText(left, top+i, line, termloop.ColorBlack, termloop.ColorRed))
	}
}
