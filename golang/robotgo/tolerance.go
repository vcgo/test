package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	toleranceArr := []float64{0.0, 0.05, 0.1, 0.15, 0.2, 0.5}
	widthArr := []int{10, 20}
	whereBitmap := robotgo.CaptureScreen()
	for _, width := range widthArr {
		findBitmap := robotgo.CaptureScreen(100, 100, width, width)
		for _, tolerance := range toleranceArr {
			ts := time.Now()
			for i := 0; i < 500; i++ {
				robotgo.FindBitmap(findBitmap, whereBitmap, tolerance)
			}
			te := time.Now()
			fmt.Println("tolerance:", tolerance, "bitmap width:", width, te.Sub(ts))
		}，，
		robotgo.FreeBitmap(findBitmap)----------------------------------------------------------------------------------------------------------
	}
}
