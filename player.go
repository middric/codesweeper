package main

import "github.com/JoelOtter/termloop"

//  Player state
const (
	Alive = 1 << iota
	Dead
)

type Player struct {
	entity *termloop.Entity
	state  int
}

func NewPlayer() Player {
	player := Player{
		entity: termloop.NewEntity(width/2, height/2, 1, 1),
		state:  Alive,
	}
	player.entity.SetCell(0, 0, &termloop.Cell{Fg: termloop.ColorBlack, Ch: '⛴'})

	return player
}

// Draw func
func (player *Player) Draw(screen *termloop.Screen) {
	if player.state != Dead {
		// Keep player in bounds
		x, y := player.entity.Position()
		if x < 0 {
			x = 0
		} else if x >= width {
			x = width - 1
		}

		if y < 0 {
			y = 0
		} else if y >= height {
			y = height - 1
		}

		player.entity.SetPosition(x, y)

		// Draw player
		player.entity.Draw(screen)

		screenWidth, screenHeight := screen.Size()
		level.SetOffset((screenWidth-width)/2, (screenHeight-height)/2)
	}
}

// Tick func
func (player *Player) Tick(event termloop.Event) {
	x, y := player.entity.Position()

	// F key
	if event.Ch == 102 && player.state != Dead && !grid.cells[x][y].isRevealed {
		if grid.cells[x][y].isFlagged {
			grid.cells[x][y].isFlagged = false
			flags--
		} else {
			grid.cells[x][y].isFlagged = true
			flags++
		}
		UpdateUI()
	} else if event.Type == termloop.EventKey {
		x, y := player.entity.Position()
		switch event.Key {
		case termloop.KeyArrowRight:
			player.entity.SetPosition(x+1, y)
			break
		case termloop.KeyArrowLeft:
			player.entity.SetPosition(x-1, y)
			break
		case termloop.KeyArrowUp:
			player.entity.SetPosition(x, y-1)
			break
		case termloop.KeyArrowDown:
			player.entity.SetPosition(x, y+1)
			break
		case termloop.KeySpace:
			if len(grid.mines) == 0 {
				grid.PlaceMines(player.entity.Position())
			}
			if grid.cells[x][y].isMine {
				ShowGameOver()
				player.state = Dead
			}
			grid.RevealCells(x, y)
			UpdateUI()
		}
	}
}
