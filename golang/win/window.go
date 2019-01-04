package main

import (
	"syscall"

	"github.com/lroc/win"
	// "github.com/lxn/win"  原 repo 没有 SetWindowText
	"github.com/vcgo/kit"
)

func main() {
	hwnd := win.FindWindow(nil, syscall.StringToUTF16Ptr("地下城与勇士"))
	if hwnd < 1 {
		kit.Fmt("Client not boot")
	}
	res1 := win.ShowWindow(hwnd, 1)
	kit.Fmt("Client ", hwnd, res1)
}
