package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

var eventsStorage []hook.Event

func main() {
	fmt.Println("Start...")

	// collect key events
	go keyListen()
	time.Sleep(time.Duration(10000) * time.Millisecond)

	// output regular
	for _, v := range eventsStorage {
		key := hook.RawcodetoKeychar(v.Rawcode)
		robotgo.KeyTap(key)
		time.Sleep(time.Duration(333) * time.Millisecond)
	}
}

func keyListen() {
	s := hook.Start()
	defer hook.End()
	for ev := range s {
		if ev.Rawcode >= 0 && ev.Kind == hook.KeyUp {
			eventsStorage = append(eventsStorage, ev)
		}
	}
}
