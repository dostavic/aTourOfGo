package main

import (
	"image"
	"image/color"
)

type Image struct {
	w, h  int
	color uint8
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{i.color, i.color, i.color, 255}
}

func main() {
	m := Image{100, 100, 255}
	pic.ShowImage(m)
}
