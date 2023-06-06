package image

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/steindlj/dc-plugins/ConnectFour/game"
)

var (
	space      = 12
	width      = 48
	Background color.RGBA
	ColorP1    color.RGBA
	ColorP2    color.RGBA
	img        = image.NewRGBA(image.Rect(0, 0, 7*width+8*space, 6*width+7*space))
)

func GenerateImg(file *os.File) {
	png.Encode(file, img)
}

func ColorBackground() {
	for i := 0; i < 7*width+8*space; i++ {
		for j := 0; j < 6*width+7*space; j++ {
			img.Set(i, j, Background)
		}
	}
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			ColorField(i, j)
		}
	}
}

func ColorField(row, col int) {
	x := (col+1)*space + col*width
	y := (row+1)*space + row*width
	color := color.RGBA{Background.R, Background.G, Background.B, uint8(float64(Background.A) * 0.9)}
	if game.Grid[row][col] != 0 {
		if game.Grid[row][col] == 1 {
			color = ColorP1
		} else {
			color = ColorP2
		}
	}
	for row = 0; row < width; row++ {
		for col = 0; col < width; col++ {
			img.SetRGBA(x+row, y+col, color)
		}
	}
}
