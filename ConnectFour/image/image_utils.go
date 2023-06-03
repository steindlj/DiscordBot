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
	ColorP1    = color.RGBA{255, 0, 0, 255}
	ColorP2    = color.RGBA{255, 255, 0, 255}
)

func GenerateImg(file *os.File) {
	img := image.NewRGBA(image.Rect(0, 0, 7*width+8*space, 6*width+7*space))
	for i := 0; i < 7*width+8*space; i++ {
		for j := 0; j < 6*width+7*space; j++ {
			img.Set(i, j, Background)
		}
	}
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			ColorField(img, i, j)
		}
	}
	png.Encode(file, img)
}

func ColorField(img *image.RGBA, i, j int) {
	x := (j+1)*space + j*width
	y := (i+1)*space + i*width
	color := color.RGBA{Background.R, Background.G, Background.B, 230}
	if game.Grid[i][j] != 0 {
		if game.Grid[i][j] == 1 {
			color = ColorP1
		} else {
			color = ColorP2
		}
	}
	for i = 0; i < width; i++ {
		for j = 0; j < width; j++ {
			img.SetRGBA(x+i, y+j, color)
		}
	}
}