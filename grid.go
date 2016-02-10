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

	return
}

func (grid *Grid) generateProximity() {
	for _, mine := range grid.mines {
		leftExtreme := (mine.x > 0)
		rightExtreme := (mine.x+1 < grid.width)

		topExtreme := (mine.y > 0)
		bottomExtreme := (mine.y+1 < grid.height)

		if leftExtreme {
			grid.cells[mine.x-1][mine.y].proximity++

			if topExtreme {
				grid.cells[mine.x-1][mine.y-1].proximity++
			}
			if bottomExtreme {
				grid.cells[mine.x-1][mine.y+1].proximity++
			}
		}

		if rightExtreme {
			grid.cells[mine.x+1][mine.y].proximity++
			if topExtreme {
				grid.cells[mine.x+1][mine.y-1].proximity++
			}
			if bottomExtreme {
				grid.cells[mine.x+1][mine.y+1].proximity++
			}
		}

		if topExtreme {
			grid.cells[mine.x][mine.y-1].proximity++
		}

		if bottomExtreme {
			grid.cells[mine.x][mine.y+1].proximity++
		}
	}
}

func NewGrid(width int, height int, mines int) Grid {
	// Create cells
	grid := Grid{cells: make([][]Cell, width), width: width, height: height}
	for x := 0; x < grid.width; x++ {
		grid.cells[x] = make([]Cell, grid.height)
		for y := 0; y < grid.height; y++ {
			grid.cells[x][y] = NewCell(x, y)
		}
	}

	// Place mines
	// grid.mines = make([]coord, mines)
	rand.Seed(time.Now().Unix())
	for i := 0; i < mines; i++ {
		grid.placeRandomMine()
	}

	grid.generateProximity()

	return grid
}
