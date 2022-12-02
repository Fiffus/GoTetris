package gamelogic

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Position struct {
	X float64
	Y float64
}

type Size struct {
	Width  int
	Height int
}

type Block struct {
	Position Position
	size     Size
	Color    *ebiten.Image
}

func (block *Block) calculateScale() [2]float64 {
	return [2]float64{float64(block.size.Width) / float64(block.Color.Bounds().Max.X), float64(block.size.Height) / float64(block.Color.Bounds().Max.Y)}
}

func InitGrid(grid *[30][15]Block, windowWidth int) {
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			grid[row][col] = Block{
				Position: Position{
					X: float64(col*30 + windowWidth/2 - len(grid[row])*30/2),
					Y: float64(row * 30),
				},
				size: Size{
					Width:  30,
					Height: 30,
				},
				Color: nil,
			}
		}
	}
}

func (block *Block) Render(screen *ebiten.Image) {
	if block.Color != nil {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Scale(block.calculateScale()[0], block.calculateScale()[1])
		options.GeoM.Translate(float64(block.Position.X), float64(block.Position.Y))
		screen.DrawImage(block.Color, options)
	}
}
