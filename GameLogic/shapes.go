package gamelogic

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	Shapes = map[string][4][2]int8{
		"SquareR0": {{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		"SquareR1": {{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		"SquareR2": {{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		"SquareR3": {{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		"LineR0":   {{0, 0}, {0, 1}, {0, 2}, {0, 3}},
		"LineR1":   {{-1, 0}, {0, 0}, {1, 0}, {2, 0}},
		"LineR2":   {{0, 0}, {0, 1}, {0, 2}, {0, 3}},
		"LineR3":   {{-1, 0}, {0, 0}, {1, 0}, {2, 0}},
		"LR0":      {{0, 0}, {0, 1}, {0, 2}, {1, 2}},
		"LR1":      {{-1, 1}, {0, 1}, {1, 0}, {1, 1}},
		"LR2":      {{-1, 0}, {0, 0}, {0, 1}, {0, 2}},
		"LR3":      {{-1, 0}, {-1, 1}, {0, 0}, {1, 0}},
		"ZR0":      {{0, 0}, {0, 1}, {1, 1}, {1, 2}},
		"ZR1":      {{0, 1}, {1, 0}, {1, 1}, {2, 0}},
		"ZR2":      {{0, 0}, {0, 1}, {1, 1}, {1, 2}},
		"ZR3":      {{0, 1}, {1, 0}, {1, 1}, {2, 0}},
		"TR0":      {{0, 1}, {1, 0}, {1, 1}, {1, 2}},
		"TR1":      {{0, 0}, {1, 0}, {2, 0}, {1, 1}},
		"TR2":      {{2, 1}, {1, 0}, {1, 1}, {1, 2}},
		"TR3":      {{0, 1}, {1, 0}, {1, 1}, {2, 1}},
	}
	ShapeBounds = map[string][2]int8{
		"SquareR0": {1, 1},
		"SquareR1": {1, 1},
		"SquareR2": {1, 1},
		"SquareR3": {1, 1},
		"LineR0":   {0, 3},
		"LineR1":   {2, 0},
		"LineR2":   {0, 3},
		"LineR3":   {2, 0},
		"LR0":      {1, 2},
		"LR1":      {1, 1},
		"LR2":      {0, 2},
		"LR3":      {1, 0},
		"ZR0":      {1, 2},
		"ZR1":      {2, 1},
		"ZR2":      {1, 2},
		"ZR3":      {2, 1},
		"TR0":      {1, 2},
		"TR1":      {2, 1},
		"TR2":      {2, 2},
		"TR3":      {2, 1},
	}
	ShapeDownStopper = map[string][][2]int8{
		"SquareR0": {{2, 0}, {2, 1}},
		"SquareR1": {{2, 0}, {2, 1}},
		"SquareR2": {{2, 0}, {2, 1}},
		"SquareR3": {{2, 0}, {2, 1}},
		"LineR0":   {{1, 0}, {1, 1}, {1, 2}, {1, 3}},
		"LineR1":   {{3, 0}},
		"LineR2":   {{1, 0}, {1, 1}, {1, 2}, {1, 3}},
		"LineR3":   {{3, 0}},
		"LR0":      {{1, 0}, {1, 1}, {2, 2}},
		"LR1":      {{2, 0}, {2, 1}},
		"LR2":      {{1, 0}, {1, 1}, {1, 2}},
		"LR3":      {{0, 1}, {2, 0}},
		"ZR0":      {{1, 0}, {2, 1}, {2, 2}},
		"ZR1":      {{3, 0}, {2, 1}},
		"ZR2":      {{1, 0}, {2, 1}, {2, 2}},
		"ZR3":      {{3, 0}, {2, 1}},
		"TR0":      {{2, 0}, {2, 1}, {2, 2}},
		"TR1":      {{2, 1}, {3, 0}},
		"TR2":      {{2, 2}, {3, 1}, {2, 0}},
		"TR3":      {{3, 1}, {2, 0}},
	}
	ShapeLeftStopper = map[string][][2]int8{
		"SquareR0": {{0, -1}, {1, -1}},
		"SquareR1": {{0, -1}, {1, -1}},
		"SquareR2": {{0, -1}, {1, -1}},
		"SquareR3": {{0, -1}, {1, -1}},
		"LineR0":   {{0, -1}},
		"LineR1":   {{-1, -1}, {0, -1}, {1, -1}, {2, -1}},
		"LineR2":   {{0, -1}},
		"LineR3":   {{-1, -1}, {0, -1}, {1, -1}, {2, -1}},
		"LR0":      {{0, -1}, {1, 1}},
		"LR1":      {{1, -1}, {0, 0}, {-1, 0}},
		"LR2":      {{0, -1}, {-1, -1}},
		"LR3":      {{-1, -1}, {0, -1}, {1, -1}},
		"ZR0":      {{0, -1}, {1, 0}},
		"ZR1":      {{2, -1}, {1, -1}, {0, 0}},
		"ZR2":      {{0, -1}, {1, 0}},
		"ZR3":      {{2, -1}, {1, -1}, {0, 0}},
		"TR0":      {{0, 0}, {1, -1}},
		"TR1":      {{0, -1}, {1, -1}, {2, -1}},
		"TR2":      {{2, 0}, {1, -1}},
		"TR3":      {{0, 0}, {1, -1}, {2, 0}},
	}
	ShapeRightStopper = map[string][][2]int8{
		"SquareR0": {{0, 2}, {1, 2}},
		"SquareR1": {{0, 2}, {1, 2}},
		"SquareR2": {{0, 2}, {1, 2}},
		"SquareR3": {{0, 2}, {1, 2}},
		"LineR0":   {{0, 4}},
		"LineR1":   {{-1, 1}, {0, 1}, {1, 1}, {2, 1}},
		"LineR2":   {{0, 4}},
		"LineR3":   {{-1, 1}, {0, 1}, {1, 1}, {2, 1}},
		"LR0":      {{0, 3}, {1, 3}},
		"LR1":      {{1, 2}, {0, 2}, {-1, 2}},
		"LR2":      {{-1, 1}, {0, 3}},
		"LR3":      {{-1, 2}, {0, 1}, {1, 1}},
		"ZR0":      {{0, 2}, {1, 3}},
		"ZR1":      {{0, 2}, {1, 2}, {2, 1}},
		"ZR2":      {{0, 2}, {1, 3}},
		"ZR3":      {{0, 2}, {1, 2}, {2, 1}},
		"TR0":      {{0, 2}, {1, 3}},
		"TR1":      {{2, 1}, {1, 2}, {0, 1}},
		"TR2":      {{2, 2}, {1, 3}},
		"TR3":      {{2, 2}, {1, 2}, {0, 2}},
	}
)

func spawnShape(player *Player, colors map[string]*ebiten.Image) {
	if player.CurrentShape == "nothing" {
		rand.Seed(time.Now().UnixNano())
		var chance int8 = int8(rand.Intn(100))
		if chance > 0 {
			player.CurrentShape = "SquareR0"
		}
		if chance > 20 {
			player.CurrentShape = "LineR0"
		}
		if chance > 40 {
			player.CurrentShape = "LR0"
		}
		if chance > 60 {
			player.CurrentShape = "ZR0"
		}
		if chance > 80 {
			player.CurrentShape = "TR0"
		}
		player.Rotation = 0
		colorKeys := make([]string, 0, len(colors))
		for key := range colors {
			colorKeys = append(colorKeys, key)
		}
		player.Color = colors[colorKeys[rand.Intn(len(colorKeys)-1)]]
	}
}

func stopShape(player *Player, grid *[30][15]Block) {
	var foundColoredBlocks int = 0
	for _, positions := range Shapes[player.CurrentShape] {
		if player.Row+positions[0] < int8(len(grid)-1) {
			if grid[player.Row+positions[0]][player.Col+positions[1]].Color != nil {
				foundColoredBlocks++
			}
		} else {
			for _, positions := range Shapes[player.CurrentShape] {
				ColoredBlocks = append(
					ColoredBlocks, ColoredBlock{
						position: [2]int8{player.Row + positions[0], player.Col + positions[1]},
						Color:    player.Color,
					},
				)
			}
			player.CurrentShape = "nothing"
			player.Color = nil
			player.DefaultPosition()
		}
	}
	if foundColoredBlocks > 0 {
		for _, positions := range Shapes[player.CurrentShape] {
			ColoredBlocks = append(
				ColoredBlocks, ColoredBlock{
					position: [2]int8{player.Row + positions[0], player.Col + positions[1]},
					Color:    player.Color,
				},
			)
		}
		player.CurrentShape = "nothing"
		player.Color = nil
		player.DefaultPosition()
		foundColoredBlocks = 0
	}
}

func drawShape(player *Player, grid *[30][15]Block) {
	for _, positions := range Shapes[player.CurrentShape] {
		grid[player.Row+positions[0]][player.Col+positions[1]].Color = player.Color
	}
}

func UpdateShape(player *Player, colors map[string]*ebiten.Image, grid *[30][15]Block) {
	spawnShape(player, colors)
	stopShape(player, grid)
	drawShape(player, grid)
	keepColoredBlocks(grid)
}
