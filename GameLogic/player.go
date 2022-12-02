package gamelogic

import (
	"fmt"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Row              int8
	Col              int8
	Color            *ebiten.Image
	CurrentShape     string
	Rotation         uint8
	rotationCooldown uint8
}

func (player *Player) fall(grid *[30][15]Block) {
	if player.Row+ShapeBounds[player.CurrentShape][0] < int8(len(grid)-1) {
		var collisions int8
		for _, positions := range ShapeDownStopper[player.CurrentShape] {
			if grid[player.Row+positions[0]][player.Col+positions[1]].Color != nil {
				collisions++
			}
		}
		if collisions < 1 {
			for _, positions := range Shapes[player.CurrentShape] {
				grid[player.Row+positions[0]][player.Col+positions[1]].Color = nil
			}
			player.Row++
		}
	}
}

func (player *Player) move(grid *[30][15]Block) {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && player.Col > 0 {
		var collisions int8
		for _, positions := range ShapeLeftStopper[player.CurrentShape] {
			if player.Col+int8(positions[1]) >= 0 {
				if grid[int8(player.Row)+positions[0]][int8(player.Col)+positions[1]].Color != nil {
					collisions++
				}
			}
		}
		if player.Row+ShapeBounds[player.CurrentShape][0] < int8(len(grid)-1) {
			for _, positions := range ShapeDownStopper[player.CurrentShape] {
				if grid[player.Row+positions[0]][player.Col+positions[1]].Color != nil {
					collisions++
				}
			}
		}
		if collisions < 1 {
			for _, positions := range Shapes[player.CurrentShape] {
				grid[player.Row+positions[0]][player.Col+positions[1]].Color = nil
			}
			player.Col--
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && player.Row+ShapeBounds[player.CurrentShape][0] < int8(len(grid)-1) {
		var collisions int8
		for _, positions := range ShapeDownStopper[player.CurrentShape] {
			if grid[player.Row+positions[0]][player.Col+positions[1]].Color != nil {
				collisions++
			}
		}
		if collisions < 1 {
			for _, positions := range Shapes[player.CurrentShape] {
				grid[player.Row+positions[0]][player.Col+positions[1]].Color = nil
			}
			player.Row++
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && player.Col+ShapeBounds[player.CurrentShape][1] < int8(len(grid[0])-1) {
		var collisions int8
		for _, positions := range ShapeRightStopper[player.CurrentShape] {
			if player.Col+positions[1] < int8(len(grid[player.Row])) {
				if grid[player.Row+positions[0]][player.Col+positions[1]].Color != nil {
					collisions++
				}
			}
		}
		if player.Row+ShapeBounds[player.CurrentShape][0] < int8(len(grid)-1) {
			for _, positions := range ShapeDownStopper[player.CurrentShape] {
				if grid[player.Row+positions[0]][player.Col+positions[1]].Color != nil {
					collisions++
				}
			}
		}
		if collisions < 1 {
			for _, positions := range Shapes[player.CurrentShape] {
				grid[player.Row+positions[0]][player.Col+positions[1]].Color = nil
			}
			player.Col++
		}
	}
}

func (player *Player) DefaultPosition() {
	player.Row = 0
	player.Col = 7
}

func (player *Player) Update(grid *[30][15]Block) {
	player.fall(grid)
	player.rotate(grid)
	player.move(grid)
}

func (player *Player) rotate(grid *[30][15]Block) {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		player.rotationCooldown++
		if player.rotationCooldown == 2 {
			var validPositions int8
			if player.Rotation+1 < 4 {
				for _, position := range Shapes[fmt.Sprintf("%vR%d", strings.Split(player.CurrentShape, "R")[0], player.Rotation+1)] {
					if player.Row+position[0] < int8(len(grid)) && player.Row+position[0] > 0 {
						if player.Col+position[1] < int8(len(grid[0])) && player.Col+position[1] > 0 {
							if grid[player.Row+position[0]][player.Col+position[1]].Color == nil {
								validPositions++
							}
						}
					}
				}
				if validPositions == int8(len(Shapes[fmt.Sprintf("%vR%d", strings.Split(player.CurrentShape, "R")[0], player.Rotation+1)])) {
					for _, position := range Shapes[player.CurrentShape] {
						grid[player.Row+position[0]][player.Col+position[1]].Color = nil
					}
					player.Rotation++
				}
			} else {
				for _, position := range Shapes[player.CurrentShape] {
					grid[player.Row+position[0]][player.Col+position[1]].Color = nil
				}
				player.Rotation = 0
			}
			if player.CurrentShape != "nothing" {
				player.CurrentShape = fmt.Sprintf("%vR%d", strings.Split(player.CurrentShape, "R")[0], player.Rotation)
			}
			player.rotationCooldown = 0
		}
	}

}

func (player *Player) CheckLose(grid *[30][15]Block, lost *bool) {
	for col := 0; col < len(grid[0]); col++ {
		if grid[2][col].Color != nil && player.Row < 1 {
			*lost = true
			return
		}
	}
	*lost = false
	return
}
