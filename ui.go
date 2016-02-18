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
		"Game Over!",
	}

	level.AddEntity(NewDialog(5, 1, msg, termloop.ColorBlack, termloop.ColorRed))
}

func GameSetup() {
	msg := []string{
		"Number of mines:",
		"",
		"{{mineCount," + strconv.Itoa(mineCount) + "}}",
	}

	level.AddEntity(NewDialog(1, 1, msg, termloop.ColorBlack, termloop.ColorGreen))
	player.state = Dead
}
