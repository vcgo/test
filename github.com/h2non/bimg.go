package main

import (
	"fmt"
	"os"

	"github.com/h2non/bimg"
)

func main() {
	buffer, err := bimg.Read("tmp/koala.png")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	watermark := bimg.Watermark{
		Text:       "Chuck Norris (c) 2315",
		Opacity:    0.25,
		Width:      200,
		DPI:        100,
		Margin:     150,
		Font:       "sans bold 12",
		Background: bimg.Color{255, 255, 255},
	}

	newImage, err := bimg.NewImage(buffer).Watermark(watermark)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	bimg.Write("tmp/watermark.jpg", newImage)
}
