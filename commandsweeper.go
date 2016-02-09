package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/JoelOtter/termloop"
)

// Play area config
const width = 50
const height = 20
const mineCount = 100

// "Graphics"
const sMine = "✱"
const sWave = "~"
const sBubble = "·"
const sSpace = " "

var game *termloop.Game
var level *termloop.BaseLevel
var grid [][]Cell
var player Player
var flags int

func updateUI() {
	screenW, screenH := game.Screen().Size()
	game.Screen().AddEntity(termloop.NewText(screenW-screenW, screenH-1, " Mines: "+strconv.Itoa(mineCount)+", Flags: "+strconv.Itoa(flags), termloop.ColorBlue, termloop.ColorBlack))
}

func gameOver(player *Player) {
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
	player.state = GameOver
}

func revealCells(x int, y int) {
	if x >= 0 && y >= 0 && x < width && y < height {
		if !grid[x][y].render && !grid[x][y].isMine {
			grid[x][y].render = true
			if grid[x][y].isFlagged {
				grid[x][y].isFlagged = false
				flags--
			}
			if grid[x][y].proximity < 1 {
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
	game.SetDebugOn(true)
	game.Screen().AddEntity(termloop.NewText(0, 0, " MineSweeper ", termloop.ColorBlue, termloop.ColorBlack))

	player = NewPlayer()

	level = termloop.NewBaseLevel(termloop.Cell{Bg: termloop.ColorBlack})
	grid = make([][]Cell, width)
	for i := range grid {
		grid[i] = make([]Cell, height)
	}

	// Set up waves
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			grid[i][j] = Cell{
				proximity: 0,
				isMine:    false,
				entity:    termloop.NewText(i, j, sSpace, termloop.ColorWhite, termloop.ColorCyan),
			}
			level.AddEntity(&grid[i][j])
		}
	}

	// Place mines
	grid = generateMines(width, height, mineCount)
	grid = generateProximity(grid)

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j].render {
				if grid[i][j].isMine {
					level.AddEntity(termloop.NewText(i, j, sMine, termloop.ColorRed, termloop.ColorCyan))
				}
				if grid[i][j].proximity > 0 {
					level.AddEntity(termloop.NewText(i, j, strconv.Itoa(grid[i][j].proximity), termloop.ColorBlack, termloop.ColorCyan))
				}
			}
		}
	}

	level.AddEntity(&player)
	game.Screen().SetLevel(level)
	game.Start()
	updateUI()
}

func generateProximity(grid [][]Cell) [][]Cell {
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y].isMine {
				leftExtreme := (x-1 >= 0)
				rightExtreme := (x+1 < len(grid))
				topExtreme := (y-1 >= 0)
				bottomExtreme := (y+1 < len(grid[x]))

				if leftExtreme {
					grid[x-1][y].proximity++

					if topExtreme {
						grid[x-1][y-1].proximity++
					}
					if bottomExtreme {
						grid[x-1][y+1].proximity++
					}
				}

				if rightExtreme {
					grid[x+1][y].proximity++
					if topExtreme {
						grid[x+1][y-1].proximity++
					}
					if bottomExtreme {
						grid[x+1][y+1].proximity++
					}
				}

				if topExtreme {
					grid[x][y-1].proximity++
				}

				if bottomExtreme {
					grid[x][y+1].proximity++
				}
			}
		}
	}
	return grid
}

func generateMines(x int, y int, count int) [][]Cell {
	rand.Seed(time.Now().Unix())
	for i := 0; i < count; i++ {
		placeMine(x, y, grid)
	}

	return grid
}

func placeMine(x int, y int, grid [][]Cell) [][]Cell {
	randX := rand.Intn(x)
	randY := rand.Intn(y)

	// If mine already in place try again
	if grid[randX][randY].isMine {
		return placeMine(x, y, grid)
	}

	grid[randX][randY].isMine = true

	return grid
}
