package gamelogic

import "github.com/hajimehoshi/ebiten/v2"

type ColoredBlock struct {
	position [2]int8
	Color    *ebiten.Image
}

var (
	ColoredBlocks []ColoredBlock
)

func keepColoredBlocks(grid *[30][15]Block) {
	for _, coloredBlock := range ColoredBlocks {
		grid[coloredBlock.position[0]][coloredBlock.position[1]].Color = coloredBlock.Color
	}
}

func CheckFullLines(grid *[30][15]Block) []int8 {
	var fullLines []int8
	for row := 0; row < len(grid); row++ {
		var count int8
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col].Color != nil {
				count++
			}
		}
		if count == int8(len(grid[row])) {
			fullLines = append(fullLines, int8(row))
			count = 0
		}
	}
	return fullLines
}

func ClearAndMove(grid *[30][15]Block, points *int, player *Player) {
	if player.Row < 1 {
		var fullLines []int8 = CheckFullLines(grid)
		for _, line := range fullLines {
			var col int8 = -1
			var coloredBlocksCopy []ColoredBlock = ColoredBlocks
			for currentBlock := len(coloredBlocksCopy) - 1; currentBlock >= 0; currentBlock-- {
				if coloredBlocksCopy[currentBlock].position[0] == line && col+1 < 15 {
					coloredBlocksCopy = append(coloredBlocksCopy[:currentBlock], coloredBlocksCopy[currentBlock+1:]...)
					col++
					grid[line][col].Color = nil
				}
			}
			ColoredBlocks = coloredBlocksCopy
			for currentBlock := len(ColoredBlocks) - 1; currentBlock >= 0; currentBlock-- {
				if ColoredBlocks[currentBlock].position[0] < line {
					grid[ColoredBlocks[currentBlock].position[0]][ColoredBlocks[currentBlock].position[1]].Color = nil
					coloredBlocksCopy[currentBlock] = ColoredBlock{
						position: [2]int8{ColoredBlocks[currentBlock].position[0] + 1, ColoredBlocks[currentBlock].position[1]},
						Color:    ColoredBlocks[currentBlock].Color,
					}
				}
			}
			ColoredBlocks = coloredBlocksCopy
			*points++
		}
	}
}
