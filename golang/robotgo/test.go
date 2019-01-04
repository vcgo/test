package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {

	w, h := 150, 150

	bitmap100 := robotgo.CaptureScreen(100, 100, w, h)
	robotgo.SaveBitmap(bitmap100, "bitmap100.png")

	bitmap200 := robotgo.CaptureScreen(200, 200, w, h)
	robotgo.SaveBitmap(bitmap200, "bitmap200.png")

	x, y := -1, -1

	x, y = robotgo.FindPic("bitmap100.png")
	fmt.Println("FindPic bitmap100.png ---------", x, y)

	x, y = robotgo.FindPic("bitmap200.png")
	fmt.Println("FindPic bitmap200.png ---------", x, y)

	x, y = robotgo.FindBitmap(bitmap100)
	fmt.Println("FindBitmap bitmap100 ---------", x, y)

	x, y = robotgo.FindBitmap(bitmap200)
	fmt.Println("FindBitmap bitmap200 ---------", x, y)
}
