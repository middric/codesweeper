package main

import "github.com/JoelOtter/termloop"

// Play area config
const width = 50
const height = 20
const mineCount = 100

var game *termloop.Game
var level *termloop.BaseLevel
var player Player
var grid Grid
var flags int

func main() {
	game = termloop.NewGame()
	player = NewPlayer()
	level = termloop.NewBaseLevel(termloop.Cell{Bg: termloop.ColorBlack})
	grid = NewGrid(width, height, mineCount)
	game.SetDebugOn(true)

	SetupUI()

	// Add cells to level
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			level.AddEntity(&grid.cells[x][y])
		}
	}

	// Add player to level
	level.AddEntity(&player)

	game.Screen().SetLevel(level)
	UpdateUI()
	GameSetup()
	game.Start()
}
