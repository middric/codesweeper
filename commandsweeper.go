package main

import (
	"strconv"

	"github.com/JoelOtter/termloop"
)

// Play area config
const width = 50
const height = 20
const mineCount = 100

var game *termloop.Game
var level *termloop.BaseLevel
var grid Grid
var player Player
var flags int

func revealCells(x int, y int) {
	if x >= 0 && y >= 0 && x < width && y < height {
		if !grid.cells[x][y].render && !grid.cells[x][y].isMine {
			grid.cells[x][y].render = true
			if grid.cells[x][y].isFlagged {
				grid.cells[x][y].isFlagged = false
				flags--
			}
			if grid.cells[x][y].proximity < 1 {
				revealCells(x-1, y-1)
				revealCells(x-1, y)
				revealCells(x+1, y+1)
				revealCells(x+1, y)
				revealCells(x, y-1)
				revealCells(x, y+1)
			}
		}
	}
}

func main() {
	game = termloop.NewGame()
	player = NewPlayer()
	level = termloop.NewBaseLevel(termloop.Cell{Bg: termloop.ColorBlack})
	grid = NewGrid(width, height, mineCount)
	game.SetDebugOn(true)

	SetupUI()

	// Set up waves
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			level.AddEntity(&grid.cells[i][j])
		}
	}

	for i := range grid.cells {
		for j := range grid.cells[i] {
			if grid.cells[i][j].render {
				if grid.cells[i][j].isMine {
					level.AddEntity(termloop.NewText(i, j, "✱", termloop.ColorRed, termloop.ColorCyan))
				}
				if grid.cells[i][j].proximity > 0 {
					level.AddEntity(termloop.NewText(i, j, strconv.Itoa(grid.cells[i][j].proximity), termloop.ColorBlack, termloop.ColorCyan))
				}
			}
		}
	}

	level.AddEntity(&player)
	game.Screen().SetLevel(level)
	game.Start()
	UpdateUI()
}
