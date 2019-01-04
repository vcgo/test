package main

import (
	"fmt"
	"reflect"

	"github.com/go-vgo/robotgo"
)

func main() {

	bitmap := robotgo.CaptureScreen(740, 140, 100, 50)
	fmt.Println("CaptureScreen...", bitmap)

	robotgo.SaveBitmap(bitmap, "double.png")

	doubleImgBit := robotgo.OpenBitmap("double.png")
	fmt.Println("openBitmap...", doubleImgBit)

	fx, fy := robotgo.FindBitmap(doubleImgBit, robotgo.CaptureScreen(), 0.01)
	fmt.Println("FindBitmap------", fx, fy)
	robotgo.Convert("test.png", "test.tif")
	bitmap := robotgo.OpenBitmap("test.tif")
	fmt.Println("...bitmap", bitmap, reflect.TypeOf(bitmap))

}
