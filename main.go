package main

import (
	"fmt"
	"image"
)

func main() {
  //here image.React represents a 100x100px image
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.ColorModel())
	fmt.Println(m.At(0, 0).RGBA())
}
