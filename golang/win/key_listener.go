// github.com/lxn/win
// win.GetKeyState 监听键盘事件
// 从而获取用户键盘操作

package main

import (
	"fmt"

	"github.com/vcgo/kit"

	"github.com/lxn/win"
)

func main() {

	hwnd := win.GetActiveWindow()
	var msg win.MSG
	win.GetMessage(&msg, hwnd, win.WM_KEYFIRST, win.WM_KEYLAST)
	fmt.Println("...", hwnd)
	return

	kl := make(chan int)

	go func() {
		for {
			kd := win.GetKeyState(win.VK_SPACE)
			if kd < 0 {
				kl <- win.VK_SPACE
				kit.Fmt("kd", kd)
			}
			kit.Sleep(20)
		}
	}()

	go func() {
		for {
			kd := win.GetKeyState(win.VK_LSHIFT)
			if kd < 0 {
				kl <- win.VK_LSHIFT
				kit.Fmt("kd", kd)
			}
			kit.Sleep(20)
		}
	}()

	var k int
	for {
		select {
		case k = <-kl:
			fmt.Println("Key press", k)
		default:
		}
	}
}
