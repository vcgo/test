package main

import (
	"github.com/vcgo/kit"
)

func main() {
	for {
		kit.LeftClick()
		kit.Sleep(233)
	}
	// Clt0
	for {
		kit.MoveClick(196, 472)
		kit.Sleep(333)
		kit.MoveClick(380, 412)
		kit.Sleep(333)
	}
	// Clt1
	for {
		kit.MoveClick(995, 466)
		kit.Sleep(333)
		kit.MoveClick(1170, 412)
		kit.Sleep(333)
	}
}
