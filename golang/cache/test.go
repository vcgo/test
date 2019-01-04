package main

import (
	"syscall"

	"github.com/lxn/win"
	. "github.com/vcgo/filecache"
	"github.com/vcgo/kit"
)

type Client struct {
	Hwnd   win.HWND
	Pid    int32
	Screen kit.Area
}

func main() {
	hwnd := win.FindWindow(nil, syscall.StringToUTF16Ptr("WeGame"))
	var rect win.RECT
	win.GetWindowRect(hwnd, &rect)
	screen := kit.Area{
		X: int(rect.Left),
		Y: int(rect.Top),
		W: int(rect.Right - rect.Left),
		H: int(rect.Bottom - rect.Top),
	}
	Clt := Client{hwnd, int32(hwnd), screen}
	kit.Fmt("clt", Clt)
	Cache.Put("clt.X", screen.X, 86400)
	Cache.Put("clt.Y", screen.Y, 86400)
	Cache.Put("clt.W", screen.W, 86400)
	Cache.Put("clt.H", screen.H, 86400)
	for i := 0; i < 10; i++ {
		kit.Fmt("clt.X", Cache.Get("clt.X").(int))
		kit.Fmt("clt.Y", Cache.Get("clt.Y"))
		kit.Fmt("clt.W", Cache.Get("clt.W"))
		kit.Fmt("clt.H", Cache.Get("clt.H"))
		kit.Sleep(1000)
	}
}
