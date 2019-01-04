package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

// https://github.com/go-vgo/robotgo/issues/133
// 若不 FreeBitmap()，内存会溢出

func main() {
	for {
		abitMap := robotgo.CaptureScreen()
		fmt.Println("abitMap...", abitMap)
		robotgo.FreeBitmap(abitMap)
	}
}
