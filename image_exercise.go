package main

import(
"golang.org/x/tour/pic"
"image"
"image/color"
)

// Define your own Image type, implement the necessary methods, and call pic.ShowImage
// source: https://go.dev/tour/methods/25
// Solution: https://gist.github.com/abhay-khandwekar/e4f2764456ed2d07d10f1e19b0fe6d70


type Image struct {
	w int
	h int
	v uint8
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{i.v, i.v, 255, 255}
}




func main() {
	m := Image{
		w: 210,
		h: 160,
		v: 143,
	}
	pic.ShowImage(m)
}
