package main

import (
	"fmt"
	"image"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
)

// 给一个图片添加文字水印
func main() {
	text := "Hello,....... world!"
	h := 100.0
	length := float64(len([]rune(text)))

	dc := gg.NewContext(int((h*0.39)*length), int(h))
	// dc.SetRGB(0, 0, 0)
	// dc.Clear()
	dc.SetRGB(1, 1, 1)
	if err := dc.LoadFontFace("/Library/Fonts/Songti.ttc", h*0.8); err != nil {
		panic(err)
	}
	fmt.Println(text)
	dc.DrawStringAnchored(text, 3, h*0.8, 0, 0)
	watermarkImage := dc.Image()

	baseImage, _ := imaging.Open("./tmp/koala.png")
	finalImage := imaging.Overlay(baseImage, watermarkImage, image.Pt(50, 50), 1.0)

	fileName := "tmp/23333333.bmp"
	imaging.Save(finalImage, fileName)

}
