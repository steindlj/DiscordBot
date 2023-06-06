package image

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/steindlj/dc-plugins/ConnectFour/game"
)

var (
	space      = 12 // between chips
	width      = 48 // width of chips 
	Background color.RGBA
	ColorP1    color.RGBA
	ColorP2    color.RGBA
	img        = image.NewRGBA(image.Rect(0, 0, 7*width+8*space, 6*width+7*space)) // holds the RGBA values set by methods
)

// Writes current image to given png-file.
func EncodeImage(file *os.File) {
	png.Encode(file, img)
}

// Generates the image by coloring the background with the background color
// and coloring each of the 42 fields based on the value in the grid. 
func GenerateImg() {
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

// Colors a specified spot in the grid based on the value in it.
// val = 0 -> no chip -> background color with 90% alpha value
// val = 1 -> player 1 -> ColorP1, val = 2 -> player 2 -> ColorP2
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
