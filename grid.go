package main

import (
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

func (grid *Grid) placeRandomMine() {
	x := rand.Intn(grid.width)
	y := rand.Intn(grid.height)

	// If mine already in place try again
	if grid.cells[x][y].isMine {
		grid.placeRandomMine()
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

func NewGrid(width int, height int, mineCount int) Grid {
	// Create cells
	grid := Grid{cells: make([][]Cell, width), width: width, height: height}
	for x := 0; x < grid.width; x++ {
		grid.cells[x] = make([]Cell, grid.height)
		for y := 0; y < grid.height; y++ {
			grid.cells[x][y] = NewCell(x, y)
		}
	}

	// Place mines
	rand.Seed(time.Now().Unix())
	for i := 0; i < mineCount; i++ {
		grid.placeRandomMine()
	}

	return grid
}
