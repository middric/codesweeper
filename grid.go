package main

import (
	"math"
	"math/rand"
	"time"
)

type coord struct {
	x int
	y int
}

type Grid struct {
	cells  [][]Cell
	mines  []coord
	width  int
	height int
}

func (grid *Grid) placeRandomMine(playerX int, playerY int) {
	x := rand.Intn(grid.width)
	y := rand.Intn(grid.height)

	// If mine already in place OR its within 2 spaces of the player try again
	if grid.cells[x][y].isMine || (math.Abs(float64(x-playerX)) <= 1 && math.Abs(float64(y-playerY)) <= 1) {
		grid.placeRandomMine(playerX, playerY)
		return
	}

	grid.cells[x][y].isMine = true
	grid.mines = append(grid.mines, coord{x, y})
	grid.updateProximity(x, y)

	return
}

func (grid *Grid) updateProximity(x int, y int) {
	if x > 0 && y > 0 {
		grid.cells[x-1][y-1].proximity++
	}
	if y > 0 {
		grid.cells[x][y-1].proximity++
	}
	if x+1 < grid.width && y > 0 {
		grid.cells[x+1][y-1].proximity++
	}
	if x > 0 {
		grid.cells[x-1][y].proximity++
	}
	if x+1 < grid.width {
		grid.cells[x+1][y].proximity++
	}
	if x > 0 && y+1 < grid.height {
		grid.cells[x-1][y+1].proximity++
	}
	if y+1 < grid.height {
		grid.cells[x][y+1].proximity++
	}
	if x+1 < grid.width && y+1 < grid.height {
		grid.cells[x+1][y+1].proximity++
	}
}

func (grid Grid) RevealCells(x int, y int) {
	if x >= 0 && y >= 0 && x < width && y < height {
		if !grid.cells[x][y].isRevealed && !grid.cells[x][y].isMine {
			grid.cells[x][y].isRevealed = true
			if grid.cells[x][y].isFlagged {
				grid.cells[x][y].isFlagged = false
				flags--
			}
			if grid.cells[x][y].proximity < 1 {
				grid.RevealCells(x-1, y-1)
				grid.RevealCells(x-1, y)
				grid.RevealCells(x+1, y+1)
				grid.RevealCells(x+1, y)
				grid.RevealCells(x, y-1)
				grid.RevealCells(x, y+1)
			}
		}
	}
}

func (grid *Grid) PlaceMines(x int, y int) {
	// Place mines
	rand.Seed(time.Now().Unix())
	for i := 0; i < mineCount; i++ {
		grid.placeRandomMine(x, y)
	}
}

func NewGrid(width int, height int, mineCount int) Grid {
	// Create cells
	grid := Grid{cells: make([][]Cell, width), width: width, height: height}
	for x := 0; x < grid.width; x++ {
		grid.cells[x] = make([]Cell, grid.height)
		for y := 0; y < grid.height; y++ {
			grid.cells[x][y] = NewCell(x, y)
		}
	}

	return grid
}
