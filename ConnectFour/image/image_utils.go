package image

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/steindlj/dc-plugins/ConnectFour/game"
)

var (
	space   = 12 // between chips
	width   = 48 // width of chips
	Grid int64
	ColorP1 int64
	ColorP2 int64
	Cell int64
	img     = image.NewRGBA(image.Rect(0, 0, 7*width+8*space, 6*width+7*space)) // holds the RGBA values set by methods
)

// EncodeImage writes the current image to the given PNG-file.
func EncodeImage(file *os.File) {
	png.Encode(file, img)
}

// GenerateImg generates the image by coloring the background with the background color
// and coloring each of the 42 fields based on the value in the grid.
func GenerateImg() {
	for i := 0; i < 7*width+8*space; i++ {
		for j := 0; j < 6*width+7*space; j++ {
			img.Set(i, j, IntToColor(Grid))
		}
	}
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			ColorField(i, j)
		}
	}
}

// ColorField colors a specified spot in the grid based on the value in it.
// val = 0 -> no chip -> Cell
// val = 1 -> player 1 -> ColorP1, val = 2 -> player 2 -> ColorP2
func ColorField(row, col int) {
	x := (col+1)*space + col*width
	y := (row+1)*space + row*width
	color := Cell
	if game.Grid[row][col] != 0 {
		if game.Grid[row][col] == 1 {
			color = ColorP1
		} else {
			color = ColorP2
		}
	}
	for row = 0; row < width; row++ {
		for col = 0; col < width; col++ {
			img.SetRGBA(x+row, y+col, IntToColor(color))
		}
	}
}

// IntToColor converts an integer value to a color.RGBA
func IntToColor(intValue int64) color.RGBA {
	red := uint8((intValue >> 16) & 0xFF)
	green := uint8((intValue >> 8) & 0xFF)
	blue := uint8(intValue & 0xFF)
	alpha := uint8(255)
	return color.RGBA{red, green, blue, alpha}
}
