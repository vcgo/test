package main

import (
	"fmt"

	"github.com/vcgo/kit"

	"github.com/go-vgo/robotgo"
)

func main() {
	for {
		bitmap := robotgo.CaptureScreen()
		goBitmap := robotgo.ToBitmap(bitmap)
		res := robotgo.SaveBitmap(bitmap, "tmp/koala.png")
		fmt.Println("CaptureScreen...", goBitmap, goBitmap.Width, res)
		robotgo.FreeBitmap(bitmap)
		kit.Sleep(99)
	}
}
