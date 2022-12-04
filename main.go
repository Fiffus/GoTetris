package main

import (
	"fmt"
	"image"
	"image/color"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"

	gamelogic "github.com/Fiffus/GoTetris/GameLogic"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
)

type Game struct{}

var (
	windowWidth  int = 700
	windowHeight int = 1015
	points       int = 0
	grid         [30][15]gamelogic.Block
	clrs         map[string]*ebiten.Image = make(map[string]*ebiten.Image)
	player       gamelogic.Player         = gamelogic.Player{
		Row:          0,
		Col:          7,
		Color:        nil,
		CurrentShape: "nothing",
	}
	backgroundImage *ebiten.Image
	backgroundSize  [2]float64
	lost            bool
	textPosition    [2]uint16
	pointsText      string
	ttf             *sfnt.Font
	textFont        font.Face
)

func init() {
	var err error
	var clrDir []fs.FileInfo
	if clrDir, err = ioutil.ReadDir("images/colors"); err != nil {
		log.Fatal(err)
	}
	for _, clr := range clrDir {
		if clrs[strings.Split(clr.Name(), ".")[0]], _, err = ebitenutil.NewImageFromFile("images/colors/" + clr.Name()); err != nil {
			log.Fatal(err)
		}
	}
	if backgroundImage, _, err = ebitenutil.NewImageFromFile("images/background/darkBackground.png"); err != nil {
		log.Fatal(err)
	}
	if ttf, err = opentype.Parse(fonts.MPlus1pRegular_ttf); err != nil {
		log.Fatal(err)
	}
	textFont, err = opentype.NewFace(ttf, &opentype.FaceOptions{
		Size:    50.0,
		DPI:     75,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	gamelogic.InitGrid(&grid, windowWidth)
	backgroundSize = [2]float64{float64(len(grid[0])*30) / float64(backgroundImage.Bounds().Max.X), float64(len(grid)*30) / float64(backgroundImage.Bounds().Max.Y)}
}

func (game *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (game *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}
	player.CheckLose(&grid, &lost)
	if lost {
		pointsText = "You lost!"
	} else {
		pointsText = fmt.Sprintf("Points: %d", points)
		player.Update(&grid)
		gamelogic.UpdateShape(&player, clrs, &grid)
		gamelogic.ClearAndMove(&grid, &points, &player)
	}
	textPosition = [2]uint16{uint16(windowWidth/2 - len(pointsText)*25/2), uint16(windowHeight - 30)}
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{35, 35, 45, 255})
	renderBackground(screen)
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			grid[row][col].Render(screen)
		}
	}
	text.Draw(screen, pointsText, textFont, int(textPosition[0]), int(textPosition[1]), color.RGBA{220, 255, 255, 255})
}

func renderBackground(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(backgroundSize[0], backgroundSize[1])
	options.GeoM.Translate(grid[0][0].Position.X, 0)
	screen.DrawImage(backgroundImage, options)
}

func main() {
	ebiten.SetMaxTPS(10)
	ebiten.SetWindowTitle("Tetris")
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowIcon([]image.Image{clrs["lightCyan"]})
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
