package main

import (
	"fmt"
	"syscall"

	"github.com/lroc/win"
	"github.com/vcgo/kit"
)

func main() {
	// hwnd0 := win.FindWindow(nil, syscall.StringToUTF16Ptr("WeGame"))
	// win.SetWindowText(hwnd0, "WeGameMaster")
	// return
	hwnd := win.FindWindow(nil, syscall.StringToUTF16Ptr("WG"))
	var rect win.RECT
	win.GetWindowRect(hwnd, &rect)
	screen := kit.Area{
		X: int(rect.Left),
		Y: int(rect.Top),
		W: int(rect.Right - rect.Left),
		H: int(rect.Bottom - rect.Top),
	}
	fmt.Println(screen)
}
