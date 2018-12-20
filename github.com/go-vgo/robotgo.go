package main

import (
	"github.com/go-vgo/robotgo"
	"github.com/vcgo/kit"
)

func main() {

	go func() {
		i := 0
		for {
			if i%3 == 0 {
				ke := robotgo.AddEvent("k")
				if ke == 0 {
					kit.Fmt("keypress k.")
				}
			} else if i%3 == 1 {
				se := robotgo.AddEvent("s")
				if se == 0 {
					kit.Fmt("keypress s.")
				}
			} else {
				robotgo.StopEvent()
			}
			i++
			kit.Fmt("i", i)
		}
	}()

	go func() {
		robotgo.AddEvent("[")
		for {
			robotgo.StopEvent()
			kit.Fmt("Wait stop event.")
			kit.Sleep(23)
		}
	}()

	for {
		kit.Sleep(9999)
	}
}
